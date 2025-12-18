package handler

import (
	"fmt"
	"main/internal/dto"
	"main/internal/services/subscriptions"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

// Create godoc
//
//	@Summary		Create new subscription
//	@Description	Returns a new subscription object.
//	@Tags			Subscription
//	@Accept			json
//	@Produce		json
//	@Param			subscription	body		dto.CreateSubRequest	true	"Subscription create data"
//	@Success		200				{object}	dto.CreateSubResponce
//	@Failure		400				{object}	handler.ErrorBadRequest
//	@Failure		500				{object}	handler.ErrorInternalError
//	@Router			/subscription [post]
func (h *handler) Create(c *gin.Context) {
	newSub := dto.CreateSubRequest{}
	err := c.BindJSON(&newSub)
	if err != nil {
		sendBadRequest(c, "request body err")
		logrus.Warn("handler create sub err:", err)
		return
	}

	id, err := h.subService.Create(c.Request.Context(), newSub)
	if err != nil {
		if err == subscriptions.ErrIncorrectDate {
			sendBadRequest(c, fmt.Sprintln(err))
			return
		}
		if err == subscriptions.ErrEndIsLess {
			sendBadRequest(c, "end date is less than start date")
			return
		}
		sendInternalError(c, "create sub err")
		return
	}

	resp := dto.CreateSubResponce{
		SubscriptionId: id,
		Success:        true,
	}
	c.JSON(http.StatusOK, resp)
}

// Read godoc
//
//	@Summary		Read subscription by ID
//	@Description	Returns a subscription object.
//	@Tags			Subscription
//	@Produce		json
//	@Param			id	path		int	true	"Subscription ID"
//	@Success		200	{object}	dto.LoadSubResponce
//	@Failure		400	{object}	handler.ErrorBadRequest
//	@Failure		404	{object}	handler.ErrorNotFound
//	@Failure		500	{object}	handler.ErrorInternalError
//	@Router			/subscription/{id} [get]
func (h *handler) Load(c *gin.Context) {
	id, err := getID(c)
	if err != nil {
		sendBadRequest(c, "sub id required")
		logrus.Warn("handler load sub err:", err)
		return
	}

	resp, err := h.subService.Load(c.Request.Context(), id)
	if err != nil {
		if err == pgx.ErrNoRows {
			sendNotFound(c, "sub not found")
			return
		}
		sendInternalError(c, "load sub err")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// List godoc
//
//	@Summary		Read subscription list
//	@Description	Returns a list of subscription objects
//	@Tags			Subscription
//	@Produce		json
//	@Param			offset	query		string	true	"offset"
//	@Param			limit	query		string	true	"limit"
//	@Success		200		{array}		dto.LoadSubResponce
//	@Failure		400		{object}	handler.ErrorBadRequest
//	@Failure		404		{object}	handler.ErrorNotFound
//	@Failure		500		{object}	handler.ErrorInternalError
//	@Router			/subscription [get]
func (h *handler) LoadList(c *gin.Context) {
	offsetStr := c.Query("offset")
	limitStr := c.Query("limit")

	offset, err := convertToInt(offsetStr)
	if err != nil {
		logrus.Warn("handler loadlist err: params invalid offset value")
		sendBadRequest(c, "params invalid offset value")
		return
	}

	limit, err := convertToInt(limitStr)
	if err != nil {
		logrus.Warn("handler loadlist err: params invalid limit value")
		sendBadRequest(c, "params invalid limit value")
		return
	}

	if limit < 0 || offset < 0 {
		logrus.Warn("handler loadlist err: limit or offset is less than 0")
		sendBadRequest(c, "limit or offset is less than 0")
		return
	}

	resp, err := h.subService.LoadList(c.Request.Context(), limit, offset)
	if err != nil {
		if err == pgx.ErrNoRows {
			sendNotFound(c, "sub list is empty")
			return
		}
		sendInternalError(c, "load sub list err")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Update godoc
//
//	@Summary		Update subscription by ID
//	@Description	Returns an ID updated subscription.
//	@Tags			Subscription
//	@Accept			json
//	@Produce		json
//	@Param			subscription	body		dto.UpdateSubRequest	true	"Subscription update data"
//	@Success		200				{object}	dto.UpdateSubResponce
//	@Failure		400				{object}	handler.ErrorBadRequest
//	@Failure		404				{object}	handler.ErrorNotFound
//	@Failure		500				{object}	handler.ErrorInternalError
//	@Router			/subscription [patch]
func (h *handler) Update(c *gin.Context) {
	req := dto.UpdateSubRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		sendBadRequest(c, "request body err")
		logrus.Warn("handler update sub err:", err)
		return
	}
	err = h.subService.Update(c.Request.Context(), req)
	if err != nil {
		if err == subscriptions.ErrIncorrectDate {
			sendBadRequest(c, fmt.Sprintln(err))
			return
		}
		if err == subscriptions.ErrEndIsLess {
			sendBadRequest(c, "end date is less than start date")
			return
		}
		if err == pgx.ErrNoRows {
			sendNotFound(c, "sub not found")
			return
		}
		sendInternalError(c, "update sub err")
		return
	}
	resp := dto.UpdateSubResponce{
		Success: true,
	}
	c.JSON(http.StatusOK, resp)
}

// Delete godoc
//
//	@Summary		Delete subscription by ID
//	@Description	Returns an ID deleted subscription.
//	@Tags			Subscription
//	@Produce		json
//	@Param			id	path		int	true	"Subscription ID"
//	@Success		200	{object}	dto.DeleteSubResponce
//	@Failure		400	{object}	handler.ErrorBadRequest
//	@Failure		404	{object}	handler.ErrorNotFound
//	@Failure		500	{object}	handler.ErrorInternalError
//	@Router			/subscription/{id} [delete]
func (h *handler) Delete(c *gin.Context) {
	id, err := getID(c)
	if err != nil {
		sendBadRequest(c, "sub id required")
		logrus.Warn("handler delete sub err:", err)
		return
	}
	err = h.subService.Delete(c.Request.Context(), id)
	if err != nil {
		if err == pgx.ErrNoRows {
			sendNotFound(c, "sub not found")
			return
		}
		sendInternalError(c, "delete sub err")
		return
	}
	resp := dto.DeleteSubResponce{
		Success: true,
	}
	c.JSON(http.StatusOK, resp)
}

// Cost godoc
//
//	@Summary		Cost subscription
//	@Description	Returns a cost of subscriptions by user ID, date and service name
//	@Tags			Subscription
//	@Accept			json
//	@Produce		json
//	@Param			subscription	body		dto.CostRequest	true	"Subscription update data"
//	@Success		200				{object}	dto.CostResponce
//	@Failure		400				{object}	handler.ErrorBadRequest
//	@Failure		404				{object}	handler.ErrorNotFound
//	@Failure		500				{object}	handler.ErrorInternalError
//	@Router			/subscription/cost [post]
func (h *handler) Cost(c *gin.Context) {
	req := dto.CostRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		sendBadRequest(c, "request body err")
		logrus.Warn("handler cost sub err:", err)
		return
	}

	resp, err := h.subService.Cost(c.Request.Context(), req)
	if err != nil {
		if err == subscriptions.ErrIncorrectDate {
			sendBadRequest(c, fmt.Sprintln(err))
			return
		}
		if err == subscriptions.ErrEndIsLess {
			sendBadRequest(c, "end date is less than start date")
			return
		}
		if err == pgx.ErrNoRows {
			sendNotFound(c, "sub not found")
			return
		}
		sendInternalError(c, "cost sub err")
		return
	}

	c.JSON(http.StatusOK, resp)
}

func convertToInt(str string) (int, error) {
	if str == "" {
		return 0, subscriptions.ErrIncorrectValue
	}

	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}

	return int(num), nil
}
