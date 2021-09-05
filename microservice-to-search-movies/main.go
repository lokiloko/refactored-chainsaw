package main

import (
	"github.com/labstack/echo/v4"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/application"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/config"
)

func main() {
	app := application.App{
		E: echo.New(),
	}

	app.InitializeController()
	app.InitializeRoutes()
	app.Start(":" + config.GetConfig().AppPort)
}
