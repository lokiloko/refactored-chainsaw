package handler

import (
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/model/dto"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/model/omdb"
)

func (h handler) GetMoviesPaginated(page uint64, keyword string) (dto.GetMoviesPaginatedResponse, error) {
	res, err := h.Service.OMDB.GetPaginated(omdb.GetPaginatedRequest{
		Keyword: keyword,
		Page:    page,
	})
	if err != nil {
		go h.Service.Logs.WriteLog(map[string]interface{}{"page": page, "keyword": keyword}, err)
		return dto.GetMoviesPaginatedResponse{}, err
	}

	go h.Service.Logs.WriteLog(map[string]interface{}{"page": page, "keyword": keyword}, res)
	return res.ToGetMoviesPaginatedResponse(), nil
}
