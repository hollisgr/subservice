package handler

import (
	"main/docs"
	"main/internal/config"
)

func initSwagger() {
	cfg := config.GetConfig()
	docs.SwaggerInfo.Title = "Subscription API server"
	docs.SwaggerInfo.Description = "This is a sample CRUDL subscription server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = cfg.Listen.Addr
	docs.SwaggerInfo.BasePath = "/"
}
