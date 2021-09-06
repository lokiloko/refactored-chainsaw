package logs

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type LogsSuite struct {
	suite.Suite
	*require.Assertions

	ctrl *gomock.Controller

	service LogsService
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(LogsSuite))
}

func (cs *LogsSuite) SetupTest() {
	cs.Assertions = require.New(cs.T())

	cs.ctrl = gomock.NewController(cs.T())

	cs.service = NewService()
}

func (cs *LogsSuite) TearDownTest() {
	cs.ctrl.Finish()
}

func (cs *LogsSuite) Test_NewService() {
	cs.Run("when create service", func() {
		c := NewService()

		cs.Equal(&logsService{}, c)
	})
}
