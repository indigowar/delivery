package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/indigowar/delivery/internal/usecases/accounts"
	"github.com/labstack/echo/v4"
)

type Delivery struct {
	finder         accounts.Finder
	registrator    accounts.Registrator
	validator      accounts.CredentialsValidator
	profileUpdater accounts.ProfileUpdater

	router *echo.Echo
	server *http.Server
}

func (d *Delivery) Run() error {
	return d.server.ListenAndServe()
}

func (d *Delivery) Shutdown(ctx context.Context) error {
	return d.server.Shutdown(ctx)
}

func (d *Delivery) AddFinder(finder accounts.Finder) {
	d.finder = finder

	// todo: setup delivery for Finder
}

func (d *Delivery) AddRegistrator(registrator accounts.Registrator) {
	d.registrator = registrator

	// todo: setup delivery for Registrator
}

func (d *Delivery) AddCredentialsValidator(validator accounts.CredentialsValidator) {
	d.validator = validator

	// todo: setup delivery for CredentialsValidator
}

func (d *Delivery) AddProfileUpdater(updater accounts.ProfileUpdater) {
	d.profileUpdater = updater

	// todo: setup delivery for ProfileUpdater
}

func NewDelivery(port int) *Delivery {
	delivery := &Delivery{}

	delivery.router = echo.New()

	delivery.router.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "ACCOUNTS SVC")
	})

	delivery.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: delivery.router,
	}

	return delivery
}
