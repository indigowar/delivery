package minio

import (
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func Connect(host string, port int, user string, password string) (*minio.Client, error) {
	endpoint := fmt.Sprintf("%s:%d", host, port)

	return minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(user, password, ""),
		Secure: true,
	})
}
