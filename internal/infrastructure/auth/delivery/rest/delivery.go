package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/indigowar/delivery/internal/usecases/auth"
	"github.com/labstack/echo/v4"
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

	router.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "AUTH SVC")
	})

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
