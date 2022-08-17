package handler

import (
	"context"
	"fmt"

	"github.com/rAndrade360/go-link-shortener/api/handler/pb"
	"github.com/teris-io/shortid"
)

type Server struct {
	pb.UnimplementedShortUrlServer
}

func (s *Server) Short(ctx context.Context, in *pb.ShortRequest) (*pb.ShortResponse, error) {
	return &pb.ShortResponse{
		ShortUrl:    fmt.Sprintf("http://0.0.0.0:8080/%s", shortid.MustGenerate()),
		OriginalUrl: in.GetUrl(),
	}, nil
}
