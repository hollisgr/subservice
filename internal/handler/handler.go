package handler

import (
	"main/internal/interfaces"

	"github.com/gin-gonic/gin"
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

func (h *handler) Register() {}
