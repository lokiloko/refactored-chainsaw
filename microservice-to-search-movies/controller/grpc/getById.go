package grpc

import "context"

func (s *Server) GetByID(ctx context.Context, in *GetByIMDBIDRequest) (*GetByIMDBIDResponse, error) {
	response, err := s.Handler.GetByIMDBID(in.GetId())
	if err != nil {
		return nil, err
	}
	return &GetByIMDBIDResponse{
		Data: &Movie{
			Title:      response.Data.Title,
			Year:       response.Data.Year,
			Rated:      *response.Data.Rated,
			Released:   *response.Data.Released,
			Runtime:    *response.Data.Runtime,
			Genre:      *response.Data.Genre,
			Director:   *response.Data.Director,
			Writer:     *response.Data.Writer,
			Actors:     *response.Data.Actors,
			Plot:       *response.Data.Plot,
			Language:   *response.Data.Language,
			Country:    *response.Data.Country,
			Awards:     *response.Data.Awards,
			Poster:     response.Data.Poster,
			Metascore:  *response.Data.Metascore,
			ImdbID:     response.Data.ImdbID,
			ImdbVotes:  *response.Data.ImdbVotes,
			ImdbRating: *response.Data.ImdbRating,
			Type:       response.Data.Type,
			Dvd:        *response.Data.DVD,
			Production: *response.Data.Production,
			Website:    *response.Data.Website,
			BoxOffice:  *response.Data.BoxOffice,
		},
	}, nil
}

