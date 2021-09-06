package grpc

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/mocks"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/model/dto"
	"github.com/pkg/errors"
)

func (cs *ControllerSuite) Test_GetPaginated() {
	cs.Run("Success", func() {
		request := GetMoviesPaginatedRequest{
			Page:          uint64(1),
			Keyword:       "batman",
		}
		ctx := context.Background()
		handlerResult := *mocks.GetOmdbPaginated(1)
		movies := []*Movie{}
		for _, movie := range handlerResult.Search {
			movies = append(movies, &Movie{
				Title:      movie.Title,
				Year:       movie.Year,
				Poster:     movie.Poster,
				ImdbID:     movie.ImdbID,
				Type:       movie.Type,
			})
		}

		expectedResult := &GetMoviesPaginatedResponse{
			Data: movies,
		}

		cs.handler.EXPECT().GetMoviesPaginated(gomock.Any(), gomock.Any()).Return(handlerResult.ToGetMoviesPaginatedResponse(), nil).Times(1)

		result, err := cs.server.GetPaginated(ctx, &request)
		cs.Nil(err)
		cs.Equal(expectedResult, result)
	})

	cs.Run("Error", func() {
		request := GetMoviesPaginatedRequest{
			Page:          uint64(1),
			Keyword:       "batman",
		}
		ctx := context.Background()
		expectedError := errors.New("force error")

		cs.handler.EXPECT().GetMoviesPaginated(gomock.Any(), gomock.Any()).Return(dto.GetMoviesPaginatedResponse{}, expectedError).Times(1)

		result, err := cs.server.GetPaginated(ctx, &request)
		cs.Equal(expectedError, err)
		cs.Nil(result)
	})
}