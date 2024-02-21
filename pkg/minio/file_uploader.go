package minio

import (
	"bytes"
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
)

type FileUploader struct {
	client  *minio.Client
	expTime time.Duration
}

func (fu *FileUploader) UploadFile(ctx context.Context, backet string, file []byte) (*url.URL, error) {
	objectName := fmt.Sprintf("image_%d.jpg", time.Now().UnixNano())

	_, err := fu.client.PutObject(ctx, backet, objectName, bytes.NewReader(file), int64(len(file)), minio.PutObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to save image: %w", err)
	}

	reqParams := make(url.Values)

	url, err := fu.client.PresignedGetObject(ctx, backet, objectName, fu.expTime, reqParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get saved image: %w", err)
	}

	return url, nil
}

func NewFileUploader(client *minio.Client, expTime time.Duration) *FileUploader {
	return &FileUploader{
		client:  client,
		expTime: expTime,
	}
}
