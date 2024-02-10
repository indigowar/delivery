package minio

import (
	"bytes"
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"

	miniohelp "github.com/indigowar/delivery/pkg/minio"
)

type ImageStorage struct {
	client *minio.Client

	backet string
}

func (is *ImageStorage) Add(ctx context.Context, image []byte) (*url.URL, error) {
	objectName := fmt.Sprintf("image_%d.jpg", time.Now().UnixNano())

	_, err := is.client.PutObject(ctx, is.backet, objectName, bytes.NewReader(image), int64(len(image)), minio.PutObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to save image: %w", err)
	}

	expire := 100 * 365 * 24 * time.Hour

	reqParams := make(url.Values)

	url, err := is.client.PresignedGetObject(ctx, is.backet, objectName, expire, reqParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get saved image: %w", err)
	}

	return url, nil
}

func NewImageStorage(host string, port int, user string, password string) (*ImageStorage, error) {
	client, err := miniohelp.Connect(host, port, user, password)
	if err != nil {
		return nil, err
	}

	return &ImageStorage{
		client: client,
		backet: "profile_pictures",
	}, nil
}
