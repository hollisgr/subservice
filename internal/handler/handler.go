package handler

import (
	"fmt"
	"main/internal/config"
	"main/internal/interfaces"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

type handler struct {
	router     *gin.Engine
	subService interfaces.Subscriptions
}

func New(r *gin.Engine, s interfaces.Subscriptions) interfaces.Handler {
	return &handler{
		router:     r,
		subService: s,
	}
}

func (h *handler) Register() {
	initSwagger()
	cfg := config.GetConfig()
	configCORS := cors.DefaultConfig()
	configCORS.AllowOrigins = cfg.CORS.AllowOrigins
	configCORS.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	configCORS.AllowCredentials = true

	h.router.Use(cors.New(configCORS))

	h.router.POST("/subscription", h.Create)
	h.router.GET("/subscription/:id", h.Load)
	h.router.GET("/subscription", h.LoadList)
	h.router.PATCH("/subscription", h.Update)
	h.router.DELETE("/subscription/:id", h.Delete)
	h.router.GET("/subscription/cost", h.Cost)

	h.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}

type ErrorBadRequest struct {
	Success bool   `json:"success" example:"false"`
	Status  string `json:"status" example:"bad request"`
	Message string `json:"message" example:"error text"`
}

type ErrorInternalError struct {
	Success bool   `json:"success" example:"false"`
	Status  string `json:"status" example:"internal error"`
	Message string `json:"message" example:"error text"`
}

type ErrorNotFound struct {
	Success bool   `json:"success" example:"false"`
	Status  string `json:"status" example:"entity not found"`
	Message string `json:"message" example:"error text"`
}

func sendBadRequest(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, ErrorBadRequest{
		Success: false,
		Message: msg,
		Status:  "bad request",
	})
}

func sendInternalError(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorInternalError{
		Success: false,
		Message: msg,
		Status:  "internal error",
	})
}

func sendNotFound(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusNotFound, ErrorNotFound{
		Success: false,
		Message: msg,
		Status:  "entity not found",
	})
}

func getID(c *gin.Context) (id int, err error) {
	s := c.Params.ByName("id")
	subId := 0
	count, err := fmt.Sscanf(s, "%d", &subId)
	if count == 0 || err != nil {
		return subId, err
	}
	return subId, nil
}
