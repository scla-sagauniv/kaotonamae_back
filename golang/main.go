package main

import (
	"kaotonamae_back/db"
	"kaotonamae_back/migrate"
	"kaotonamae_back/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// DB接続
	db.Init()
	migrate.Run()

	server := echo.New()
	routes.RegisterRoutes(server)

	server.Logger.Fatal(server.Start(":8080"))

}
