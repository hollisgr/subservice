package main

import (
	"main/internal/app"
	"main/internal/config"
	"main/internal/db"
	"main/internal/services/subscriptions"
)

func main() {
	cfg := config.GetConfig()

	pgxPool := app.ConnectToDB(cfg)
	defer pgxPool.Close()

	app.InitLogger(cfg.Logger.LogLevel)

	storage := db.New(pgxPool)

	subServ := subscriptions.New(storage)

	router := app.SetupRouter(subServ)

	server := app.SetupServer(cfg, router)

	app.StartServer(server)
}
