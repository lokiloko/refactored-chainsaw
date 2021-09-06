package omdb

import (
	"fmt"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/config"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/model/dto"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/model/omdb"
	"go.uber.org/zap"
)

func (r omdbService) GetByIMDBID(request omdb.GetByIDRequest) (omdb.GetByIDResponse, error) {
	resp := omdb.GetByIDResponse{}
	requestParams := dto.WebServiceRequestParams{
		Method: "GET",
		Url: fmt.Sprintf("%v?apikey=%v&i=%v",
			config.GetConfig().OMDBHost, config.GetConfig().OMDBAPIKey, request.ID),
	}
	if err := Request(&requestParams, &resp); err != nil {
		zap.S().Error("get by id error ", err)
		return resp, err
	}

	return resp, nil
}
