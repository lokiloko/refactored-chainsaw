package omdb

import (
	"fmt"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/config"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/model/dto"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/model/omdb"
	"go.uber.org/zap"
)

func (r omdbService) GetPaginated(request omdb.GetPaginatedRequest) (omdb.GetPaginatedResponse, error) {
	resp := omdb.GetPaginatedResponse{}
	requestParams := dto.WebServiceRequestParams{
		Method: "GET",
		Url: fmt.Sprintf("%v?apikey=%v&s=%v&page=%v",
			config.GetConfig().OMDBHost, config.GetConfig().OMDBAPIKey, request.Keyword, request.Page),
	}
	if err := Request(&requestParams, &resp); err != nil {
		zap.S().Error("get paginated error ", err)
		return resp, err
	}

	return resp, nil
}
