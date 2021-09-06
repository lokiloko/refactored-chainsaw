package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/model/dto"
	"net/http"
	"strconv"
)

func (controller *Controller) GetMoviesPaginated(c echo.Context) error {
	var (
		response dto.GetMoviesPaginatedResponse
	)
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}
	response, err = controller.handler.GetMoviesPaginated(uint64(page), c.QueryParam("keyword"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, response)
}
