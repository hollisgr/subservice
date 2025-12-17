package handler

import (
	"fmt"
	"main/internal/config"
	"main/internal/interfaces"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

type handler struct {
	router     *gin.Engine
	subService interfaces.Subscriptions
}

func New(r *gin.Engine, s interfaces.Subscriptions) interfaces.Handler {
	initSwagger()
	return &handler{
		router:     r,
		subService: s,
	}
}

func (h *handler) Register() {
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

	h.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}

type RespMsgError struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"error text"`
}

func initSwagger() {
	cfg := config.GetConfig()
	docs.SwaggerInfo.Title = "Subscription API server"
	docs.SwaggerInfo.Description = "This is a sample CRUDL subscription server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = cfg.Listen.Addr
	docs.SwaggerInfo.BasePath = "/"
}

func sendError(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, RespMsgError{
		Success: false,
		Message: msg,
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
