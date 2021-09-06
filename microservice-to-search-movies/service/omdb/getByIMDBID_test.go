package omdb

import (
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/mocks"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/model/omdb"
)

func (s *OMDBSuite) Test_GetByID() {
	var request = omdb.GetByIDRequest{
		ID: "tt0372784",
	}
	s.Run("when success", func() {
		expectedResult := *mocks.GetOMDBByID(1)

		result, err := s.service.GetByIMDBID(request)

		s.Nil(err)
		s.Equal(expectedResult, result)
	})
}
