package handler

import (
	"github.com/golang/mock/gomock"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/mocks"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/model/dto"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/model/omdb"
	"github.com/pkg/errors"
	"sync"
)

func (s *HandlerSuite) Test_GetByIMDBId() {
	s.Run("Success", func() {
		id := "tt0372784"
		expectedResult := *mocks.GetOMDBByID(1)

		var wg sync.WaitGroup
		wg.Add(1)

		s.omdbService.EXPECT().GetByIMDBID(gomock.Any()).Return(expectedResult, nil).Times(1)
		s.logsService.EXPECT().WriteLog(gomock.Any(), gomock.Any()).Do(func(a, b interface{}) {
			wg.Done()
		}).Return(nil).Times(1)

		result, err := s.handler.GetByIMDBID(id)
		wg.Wait()
		s.Nil(err)
		s.Equal(expectedResult.ToGetByIMDBIDResponse(), result)
	})

	s.Run("Error", func() {
		id := "tt0372784"
		expectedError := errors.New("force error")

		var wg sync.WaitGroup
		wg.Add(1)

		s.omdbService.EXPECT().GetByIMDBID(gomock.Any()).Return(omdb.GetByIDResponse{}, expectedError).Times(1)
		s.logsService.EXPECT().WriteLog(gomock.Any(), gomock.Any()).Do(func(a, b interface{}) {
			wg.Done()
		}).Return(nil).Times(1)

		result, err := s.handler.GetByIMDBID(id)
		wg.Wait()
		s.Equal(expectedError, err)
		s.Equal(dto.GetByIMDBIDResponse{}, result)
	})
}
