package main

import (
	"log"
	"net/http"
	"os"

	"github.com/indigowar/delivery/internal/infrastructure/accounts/storage/postgres"
	"github.com/indigowar/delivery/internal/usecases/accounts"
)

func main() {
	_, err := createPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ACCOUNTS SVC"))
	})

	_ = http.ListenAndServe(":80", nil)
}

func createPostgresStorage() (accounts.Storage, error) {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbName := os.Getenv("POSTGRES_DB")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")

	return postgres.NewStorage(host, port, user, password, dbName)
}
