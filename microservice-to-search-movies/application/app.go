package application

import (
	"github.com/labstack/echo/v4"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/controller"
	"go.uber.org/zap"
)

type App struct {
	E          *echo.Echo
	Controller *controller.Controller
}

func (app *App) Start(address string) {
	err := app.E.Start(address)
	if err != nil {
		zap.Error(err)
		panic(err)
	}
}

func (app *App) InitializeRoutes() {
	_ = app.E.Group("/v1")

	app.E.GET("/health", app.Controller.Health)
}

func (app *App) InitializeController() {
	ctrl := controller.NewController()

	app.Controller = ctrl
}
