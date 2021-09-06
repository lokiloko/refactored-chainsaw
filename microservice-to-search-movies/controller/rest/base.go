package rest

import "github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/handler"

type Controller struct {
	handler handler.Handler
}

func NewController(handler handler.Handler) *Controller {
	return &Controller{
		handler: handler,
	}
}