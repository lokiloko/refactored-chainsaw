package grpc

import (
	"context"
)

func (s *Server) GetPaginated(ctx context.Context, request *GetMoviesPaginatedRequest) (*GetMoviesPaginatedResponse, error) {
	response, err := s.Handler.GetMoviesPaginated(request.GetPage(), request.GetKeyword())
	if err != nil {
		return nil, err
	}

	movies := []*Movie{}
	for _, movie := range response.Data {
		movies = append(movies, &Movie{
			Title:      movie.Title,
			Year:       movie.Year,
			Poster:     movie.Poster,
			ImdbID:     movie.ImdbID,
			Type:       movie.Type,
		})
	}
	return &GetMoviesPaginatedResponse{
		Data: movies,
	}, nil
}
