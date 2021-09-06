package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/model/dto"
	"net/http"
)

func (controller *Controller) GetByIMDBID(c echo.Context) error {
	var (
		response dto.GetByIMDBIDResponse
	)

	response, err := controller.handler.GetByIMDBID(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, response)
}