package application

import (
	"github.com/labstack/echo/v4"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/controller/rest"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/handler"
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
	v1 := app.E.Group("/v1")

	app.E.GET("/health", app.Controller.Health)
	v1.GET("/movie/:id", app.Controller.GetByIMDBID)
	v1.GET("/movie", app.Controller.GetMoviesPaginated)
}

func (app *App) InitializeRestController() {
	svc := service.Service{
		OMDB: omdb.NewService(),
	}
	hdlr := handler.NewHandler(svc)

	ctrl := rest.NewController(hdlr)

	app.Controller = ctrl
}
