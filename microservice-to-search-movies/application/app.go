package application

import (
	"github.com/labstack/echo/v4"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/controller/rest"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/service"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/service/omdb"
	"go.uber.org/zap"
)

type App struct {
	E          *echo.Echo
	Controller *rest.Controller
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
	ctrl := rest.NewController()
	_ = service.Service{
		OMDB: omdb.NewService(),
	}

	app.Controller = ctrl
}
