package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/indigowar/delivery/internal/usecases/auth"
	"github.com/indigowar/delivery/pkg/http/status"
)

type Delivery struct {
	service auth.Service

	router *echo.Echo
	server *http.Server
}

func (d *Delivery) Run() error {
	return d.server.ListenAndServe()
}

func (d *Delivery) Shutdown(ctx context.Context) error {
	return d.server.Shutdown(ctx)
}

func (d *Delivery) AddService(service auth.Service) {
	d.service = service

	d.router.POST("/session", startSessionHandler(d.service))
	d.router.DELETE("/session", endSessionHandler(d.service))
	d.router.PUT("/session", extendSessionHandler(d.service))
	d.router.POST("/session/refresh", getAccessTokenHandler(d.service))
}

func NewDelivery(port uint) *Delivery {
	router := echo.New()

	router.GET("/status", status.StatusHandler("auth"))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	return &Delivery{
		service: nil,
		router:  router,
		server:  server,
	}
}
