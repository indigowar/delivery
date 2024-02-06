package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/indigowar/delivery/internal/infrastructure/auth/storage/redis"
	connector "github.com/indigowar/delivery/pkg/redis"
)

func main() {
	storage, err := createRedisStorage()
	if err != nil {
		log.Fatal(err)
	}
	defer storage.Close()

	// todo: add the server

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("AUTH SERVICE"))
	})

	_ = http.ListenAndServe(":80", nil)
}

func createRedisStorage() (*redis.Storage, error) {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")

	client, err := connector.ConnectToRedis(host, port, password)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return redis.NewStorage(client), nil
}
