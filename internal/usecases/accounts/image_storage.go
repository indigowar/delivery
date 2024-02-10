package accounts

import (
	"context"
	"net/url"
)

//go:generate moq -out image_storage_moq_test.go . ImageStorage

type ImageStorage interface {
	Add(ctx context.Context, image []byte) (*url.URL, error)
}
