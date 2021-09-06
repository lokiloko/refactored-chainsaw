package grpc

import (
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/handler"
)

type Server struct {
	Handler handler.Handler
}

func (s *Server) mustEmbedUnimplementedMovieServiceServer() {
	panic("implement me")
}