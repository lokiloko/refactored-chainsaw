package rest

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ControllerSuite struct {
	suite.Suite
	*require.Assertions

	ctrl *gomock.Controller

	controller *Controller
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(ControllerSuite))
}

func (cs *ControllerSuite) SetupTest() {
	cs.Assertions = require.New(cs.T())

	cs.ctrl = gomock.NewController(cs.T())

	cs.controller = NewController()
}

func (cs *ControllerSuite) TearDownTest() {
	cs.ctrl.Finish()
}

func (cs *ControllerSuite) Test_NewController() {
	cs.Run("when create controller", func() {
		c := NewController()

		cs.Equal(&Controller{}, c)
	})
}
