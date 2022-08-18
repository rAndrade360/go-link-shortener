package adapter

import (
	"context"
	"log"

	"github.com/rAndrade360/go-link-shortener/bot/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type api struct {
	client pb.ShortUrlClient
}

type Api interface {
	ShortUrl(url string) (string, error)
}

func NewAdapterApi(address string) Api {
	con, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error to GRPC Dial", err.Error())
	}

	return &api{
		client: pb.NewShortUrlClient(con),
	}
}

func (a *api) ShortUrl(url string) (string, error) {
	shortUrl, err := a.client.Short(context.Background(), &pb.ShortRequest{
		Url: url,
	})
	if err != nil {
		return "", err
	}

	return shortUrl.ShortUrl, nil
}
