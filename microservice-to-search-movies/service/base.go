package service

import "github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/service/omdb"

type Service struct {
	OMDB omdb.OMDBService
}
