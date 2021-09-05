package application

import (
	"fmt"
	"github.com/labstack/echo/v4"
	grpc2 "github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/controller/grpc"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/controller/rest"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/handler"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/service"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/service/logs"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/service/omdb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	E          *echo.Echo
	Controller *rest.Controller
	Grpc       *grpc.Server
	Service    service.Service
	Handler    handler.Handler
}

func (app *App) Start(address string) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		zap.S().Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()
	go app.StartGrpc(lis)

	err = app.E.Start(address)
	if err != nil {
		zap.Error(err)
		panic(err)
	}
}

func (app *App) StartGrpc(lis net.Listener) {
	if err := app.Grpc.Serve(lis); err != nil {
		zap.S().Fatalf("failed to serve: %s", err)
	}
}

func (app *App) InitializeRoutes() {
	v1 := app.E.Group("/v1")

	app.E.GET("/health", app.Controller.Health)
	v1.GET("/movie/:id", app.Controller.GetByIMDBID)
	v1.GET("/movie", app.Controller.GetMoviesPaginated)
}

func (app *App) InitializeGrpc() {
	app.Grpc = grpc.NewServer()

	s := grpc2.Server{
		Handler: app.Handler,
	}
	grpc2.RegisterMovieServiceServer(app.Grpc, &s)
}

func (app *App) InitializeRestController() {
	ctrl := rest.NewController(app.Handler)

	app.Controller = ctrl
}

func (app *App) InitializeService() {
	svc := service.Service{
		OMDB: omdb.NewService(),
		Logs: logs.NewService(),
	}

	app.Service = svc
}

func (app *App) InitializeHandler() {
	hdlr := handler.NewHandler(app.Service)

	app.Handler = hdlr
}
