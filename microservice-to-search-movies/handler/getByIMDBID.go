package handler

import (
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/model/dto"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/model/omdb"
)

func (h handler) GetByIMDBID(id string) (dto.GetByIMDBIDResponse, error) {
	res, err := h.Service.OMDB.GetByIMDBID(omdb.GetByIDRequest{
		ID: id,
	})
	if err != nil {
		return dto.GetByIMDBIDResponse{}, err
	}

	return res.ToGetByIMDBIDResponse(), nil
}
