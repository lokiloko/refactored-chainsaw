package rest

import (
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/mocks"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/model/dto"
	"github.com/pkg/errors"
	"net/http"
	"net/http/httptest"
)

func (cs *ControllerSuite) Test_GetByIMDBId() {
	cs.Run("Success", func() {
		id := "tt0372784"
		expectedResult := *mocks.GetOMDBByID(1)

		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/v1/movie/:id", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		q := req.URL.Query()
		q.Add("id", id)
		req.URL.RawQuery = q.Encode()

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		expectedStatusCode := http.StatusOK

		cs.handler.EXPECT().GetByIMDBID(gomock.Any()).Return(expectedResult.ToGetByIMDBIDResponse(), nil).Times(1)

		_ = cs.controller.GetByIMDBID(c)
		cs.Equal(expectedStatusCode, rec.Code)
	})

	cs.Run("Error", func() {
		id := "tt0372784"

		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/v1/movie/:id", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		q := req.URL.Query()
		q.Add("id", id)
		req.URL.RawQuery = q.Encode()

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		expectedStatusCode := http.StatusInternalServerError

		cs.handler.EXPECT().GetByIMDBID(gomock.Any()).Return(dto.GetByIMDBIDResponse{}, errors.New("force error")).Times(1)

		_ = cs.controller.GetByIMDBID(c)
		cs.Equal(expectedStatusCode, rec.Code)
	})
}
