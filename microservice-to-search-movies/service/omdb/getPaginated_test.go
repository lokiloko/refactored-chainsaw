package omdb

import (
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/mocks"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/model/omdb"
)

func (s *OMDBSuite) Test_GetPaginated() {
	var request = omdb.GetPaginatedRequest{
		Keyword: "batman",
		Page:    1,
	}
	s.Run("when success", func() {
		expectedResult := *mocks.GetOmdbPaginated(1)

		result, err := s.service.GetPaginated(request)

		s.Nil(err)
		s.Equal(expectedResult, result)
	})
}

