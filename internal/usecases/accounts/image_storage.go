package accounts

import (
	"context"
	"errors"
	"net/url"
)

//go:generate moq -out image_storage_moq_test.go . ImageStorage

var (
	ErrImageNotFound = errors.New("image is not found")
)

type ImageStorage interface {
	Get(ctx context.Context, url *url.URL) ([]byte, error)
	Add(ctx context.Context, image []byte) (*url.URL, error)
}
