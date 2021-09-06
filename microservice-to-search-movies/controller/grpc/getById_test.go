package grpc

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/mocks"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/model/dto"
	"github.com/pkg/errors"
)

func (cs *ControllerSuite) Test_GetByIMDBId() {
	cs.Run("Success", func() {
		id := "tt0372784"
		request := GetByIMDBIDRequest{
			Id: id,
		}
		ctx := context.Background()
		handlerResult := *mocks.GetOMDBByID(1)
		expectedResult := &GetByIMDBIDResponse{
			Data: &Movie{
				Title:      handlerResult.Title,
				Year:       handlerResult.Year,
				Rated:      *handlerResult.Rated,
				Released:   *handlerResult.Released,
				Runtime:    *handlerResult.Runtime,
				Genre:      *handlerResult.Genre,
				Director:   *handlerResult.Director,
				Writer:     *handlerResult.Writer,
				Actors:     *handlerResult.Actors,
				Plot:       *handlerResult.Plot,
				Language:   *handlerResult.Language,
				Country:    *handlerResult.Country,
				Awards:     *handlerResult.Awards,
				Poster:     handlerResult.Poster,
				Metascore:  *handlerResult.Metascore,
				ImdbID:     handlerResult.ImdbID,
				ImdbVotes:  *handlerResult.ImdbVotes,
				ImdbRating: *handlerResult.ImdbRating,
				Type:       handlerResult.Type,
				Dvd:        *handlerResult.DVD,
				Production: *handlerResult.Production,
				Website:    *handlerResult.Website,
				BoxOffice:  *handlerResult.BoxOffice,
			},
		}

		cs.handler.EXPECT().GetByIMDBID(gomock.Any()).Return(handlerResult.ToGetByIMDBIDResponse(), nil).Times(1)

		result, err := cs.server.GetByID(ctx, &request)
		cs.Nil(err)
		cs.Equal(expectedResult, result)
	})

	cs.Run("Error", func() {
		id := "tt0372784"
		request := GetByIMDBIDRequest{
			Id: id,
		}
		ctx := context.Background()

		expectedError := errors.New("force error")

		cs.handler.EXPECT().GetByIMDBID(gomock.Any()).Return(dto.GetByIMDBIDResponse{}, expectedError).Times(1)

		result, err := cs.server.GetByID(ctx, &request)
		cs.Equal(expectedError, err)
		cs.Nil(result)
	})
}
