package handler

import (
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/model/dto"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/service"
)

type handler struct {
	Service service.Service
}

type Handler interface {
	GetByIMDBID(id string) (dto.GetByIMDBIDResponse, error)
	GetMoviesPaginated(page uint64, keyword string) (dto.GetMoviesPaginatedResponse, error)
}

func NewHandler(service service.Service) Handler {
	return &handler{
		Service: service,
	}
}
