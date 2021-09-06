package grpc

import (
	"github.com/golang/mock/gomock"
	mock_handler "github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/mocks/handler"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ControllerSuite struct {
	suite.Suite
	*require.Assertions

	ctrl *gomock.Controller

	handler *mock_handler.MockHandler
	server  *Server
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(ControllerSuite))
}

func (cs *ControllerSuite) SetupTest() {
	cs.Assertions = require.New(cs.T())

	cs.ctrl = gomock.NewController(cs.T())
	cs.handler = mock_handler.NewMockHandler(cs.ctrl)

	cs.server = &Server{
		Handler: cs.handler,
	}
}

func (cs *ControllerSuite) TearDownTest() {
	cs.ctrl.Finish()
}
