package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
)

func (cs *ControllerSuite) Test_Health() {
	cs.Run("Success", func() {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/health", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		expectedStatusCode := http.StatusOK

		_ = cs.controller.Health(c)
		cs.Equal(expectedStatusCode, rec.Code)
	})
}