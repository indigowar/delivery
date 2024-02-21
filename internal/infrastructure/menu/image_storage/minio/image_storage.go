package minio

import (
	"context"
	"net/url"
	"time"

	"github.com/indigowar/delivery/pkg/minio"
)

type ImageStorage struct {
	uploader *minio.FileUploader

	backet string
}

func (is *ImageStorage) Add(ctx context.Context, image []byte) (*url.URL, error) {
	return is.uploader.UploadFile(ctx, is.backet, image)
}

func NewImageStorage(host string, port int, user string, password string) (*ImageStorage, error) {
	client, err := minio.Connect(host, port, user, password)
	if err != nil {
		return nil, err
	}

	expTime := 100 * 365 * 24 * time.Hour

	return &ImageStorage{
		uploader: minio.NewFileUploader(client, expTime),
		backet:   "menu_images",
	}, nil
}
