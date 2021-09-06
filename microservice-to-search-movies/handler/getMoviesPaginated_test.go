package handler

import (
	"github.com/golang/mock/gomock"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/mocks"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/model/dto"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/model/omdb"
	"github.com/pkg/errors"
	"sync"
)

func (s *HandlerSuite) Test_GetMoviesPaginated() {
	s.Run("Success", func() {
		page := uint64(1)
		keyword := "batman"
		expectedResult := *mocks.GetOmdbPaginated(1)

		var wg sync.WaitGroup
		wg.Add(1)

		s.omdbService.EXPECT().GetPaginated(gomock.Any()).Return(expectedResult, nil).Times(1)
		s.logsService.EXPECT().WriteLog(gomock.Any(), gomock.Any()).Do(func(a, b interface{}) {
			wg.Done()
		}).Return(nil).Times(1)

		result, err := s.handler.GetMoviesPaginated(page, keyword)
		wg.Wait()
		s.Nil(err)
		s.Equal(expectedResult.ToGetMoviesPaginatedResponse(), result)
	})

	s.Run("Error", func() {
		page := uint64(1)
		keyword := "batman"
		expectedError := errors.New("force error")

		var wg sync.WaitGroup
		wg.Add(1)

		s.omdbService.EXPECT().GetPaginated(gomock.Any()).Return(omdb.GetPaginatedResponse{}, expectedError).Times(1)
		s.logsService.EXPECT().WriteLog(gomock.Any(), gomock.Any()).Do(func(a, b interface{}) {
			wg.Done()
		}).Return(nil).Times(1)

		result, err := s.handler.GetMoviesPaginated(page, keyword)
		wg.Wait()
		s.Equal(expectedError, err)
		s.Equal(dto.GetMoviesPaginatedResponse{}, result)
	})
}
