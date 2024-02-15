package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/indigowar/delivery/pkg/http/status"
)

func main() {
	router := echo.New()

	router.GET("/status", status.StatusHandler("menu"))

	server := &http.Server{
		Addr:    ":80",
		Handler: router,
	}

	_ = server.ListenAndServe()
}
