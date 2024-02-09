package minio

import (
	"context"
	"net/url"

	"github.com/minio/minio-go/v7"

	miniohelp "github.com/indigowar/delivery/pkg/minio"
)

type ImageStorage struct {
	client *minio.Client
}

func (is *ImageStorage) Get(ctx context.Context, url *url.URL) ([]byte, error) {
	panic("not implemented")
}

func (is *ImageStorage) Add(ctx context.Context, image []byte) (*url.URL, error) {
	panic("not implemented")
}

func NewImageStorage(host string, port int, user string, password string) (*ImageStorage, error) {
	client, err := miniohelp.Connect(host, port, user, password)
	if err != nil {
		return nil, err
	}

	return &ImageStorage{
		client: client,
	}, nil
}

// todo: implement
