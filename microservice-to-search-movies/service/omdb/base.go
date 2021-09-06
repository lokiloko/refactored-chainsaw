package omdb

import "github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/model/omdb"

type OMDBService interface {
	GetPaginated(request omdb.GetPaginatedRequest) (omdb.GetPaginatedResponse, error)
	GetByIMDBID(request omdb.GetByIDRequest) (omdb.GetByIDResponse, error)
}

type omdbService struct {
}

func NewOMDBService() OMDBService {
	return &omdbService{}
}
