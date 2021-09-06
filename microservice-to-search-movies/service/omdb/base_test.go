package omdb

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type OMDBSuite struct {
	suite.Suite
	*require.Assertions

	ctrl *gomock.Controller

	service OMDBService
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OMDBSuite))
}

func (cs *OMDBSuite) SetupTest() {
	cs.Assertions = require.New(cs.T())

	cs.ctrl = gomock.NewController(cs.T())

	cs.service = NewOMDBService()
}

func (cs *OMDBSuite) TearDownTest() {
	cs.ctrl.Finish()
}

func (cs *OMDBSuite) Test_NewService() {
	cs.Run("when create service", func() {
		c := NewOMDBService()

		cs.Equal(&omdbService{}, c)
	})
}