package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"

	"github.com/indigowar/delivery/pkg/http/status"
	"github.com/indigowar/delivery/pkg/postgres"
)

func main() {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbName := os.Getenv("POSTGRES_DB")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")

	_, err := postgres.Connect(host, port, user, password, dbName)
	if err != nil {
		log.Fatal(err)
	}

	router := echo.New()

	router.GET("/status", status.StatusHandler("menu"))

	server := &http.Server{
		Addr:    ":80",
		Handler: router,
	}

	_ = server.ListenAndServe()
}
