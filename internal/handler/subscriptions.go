package handler

import (
	"main/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func (h *handler) Create(c *gin.Context) {
	newSub := dto.CreateSubRequest{}
	err := c.BindJSON(&newSub)
	if err != nil {
		sendError(c, http.StatusBadRequest, "request body required")
		logrus.Warn("handler create sub err:", err)
		return
	}

	id, err := h.subService.Create(c.Request.Context(), newSub)
	if err != nil {
		sendError(c, http.StatusInternalServerError, "create sub err")
		return
	}

	resp := dto.CreateSubResponce{
		SubscriptionId: id,
		Success:        true,
	}
	c.JSON(http.StatusOK, resp)
}

func (h *handler) Load(c *gin.Context) {
	id, err := getID(c)
	if err != nil {
		sendError(c, http.StatusBadRequest, "sub id required")
		logrus.Warn("handler load sub err:", err)
		return
	}

	resp, err := h.subService.Load(c.Request.Context(), id)
	if err != nil {
		if err == pgx.ErrNoRows {
			sendError(c, http.StatusNotFound, "sub not found")
			return
		}
		sendError(c, http.StatusInternalServerError, "load sub err")
		return
	}

	c.JSON(http.StatusOK, resp)
}
func (h *handler) LoadList(c *gin.Context) {
	req := dto.LoadListRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		sendError(c, http.StatusBadRequest, "request body required")
		logrus.Warn("handler load sub list err:", err)
		return
	}
	resp, err := h.subService.LoadList(c.Request.Context(), req.Limit, req.Offset)
	if err != nil {
		if err == pgx.ErrNoRows {
			sendError(c, http.StatusNotFound, "sub list not found")
			return
		}
		sendError(c, http.StatusInternalServerError, "load sub list err")
		return
	}
	c.JSON(http.StatusOK, resp)
}
func (h *handler) Update(c *gin.Context) {
	req := dto.UpdateSubRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		sendError(c, http.StatusBadRequest, "request body required")
		logrus.Warn("handler update sub err:", err)
		return
	}
	err = h.subService.Update(c.Request.Context(), req)
	if err != nil {
		if err == pgx.ErrNoRows {
			sendError(c, http.StatusNotFound, "sub not found, or data equal")
			return
		}
		sendError(c, http.StatusInternalServerError, "update sub err")
		return
	}
	resp := dto.UpdateSubResponce{
		Success: true,
	}
	c.JSON(http.StatusOK, resp)
}
func (h *handler) Delete(c *gin.Context) {
	id, err := getID(c)
	if err != nil {
		sendError(c, http.StatusBadRequest, "sub id required")
		logrus.Warn("handler delete sub err:", err)
		return
	}
	err = h.subService.Delete(c.Request.Context(), id)
	if err != nil {
		if err == pgx.ErrNoRows {
			sendError(c, http.StatusNotFound, "sub not found")
			return
		}
		sendError(c, http.StatusInternalServerError, "delete sub err")
		return
	}
	resp := dto.DeleteSubResponce{
		Success: true,
	}
	c.JSON(http.StatusOK, resp)
}
