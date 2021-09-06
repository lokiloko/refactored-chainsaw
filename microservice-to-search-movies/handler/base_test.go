package handler

import (
	"github.com/golang/mock/gomock"
	mock_service "github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/mocks/service/logs"
	mock_service_omdb "github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/mocks/service/omdb"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/service"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type HandlerSuite struct {
	suite.Suite
	*require.Assertions

	ctrl *gomock.Controller
	service service.Service
	omdbService *mock_service_omdb.MockOMDBService
	logsService *mock_service.MockLogsService

	handler Handler
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}

func (cs *HandlerSuite) SetupTest() {
	cs.Assertions = require.New(cs.T())

	cs.ctrl = gomock.NewController(cs.T())
	cs.omdbService = mock_service_omdb.NewMockOMDBService(cs.ctrl)
	cs.logsService = mock_service.NewMockLogsService(cs.ctrl)
	cs.service = service.Service{
		OMDB: cs.omdbService,
		Logs: cs.logsService,
	}

	cs.handler = NewHandler(cs.service)
}

func (cs *HandlerSuite) TearDownTest() {
	cs.ctrl.Finish()
}

func (cs *HandlerSuite) Test_NewController() {
	cs.Run("when create controller", func() {
		svc := service.Service{
			OMDB: mock_service_omdb.NewMockOMDBService(cs.ctrl),
			Logs: mock_service.NewMockLogsService(cs.ctrl),
		}
		c := NewHandler(svc)

		cs.Equal(&handler{
			Service: svc,
		}, c)
	})
}

