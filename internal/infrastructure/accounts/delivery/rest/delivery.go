package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/indigowar/delivery/internal/infrastructure/accounts/delivery/rest/handlers"
	"github.com/indigowar/delivery/internal/infrastructure/accounts/delivery/rest/middleware"
	"github.com/indigowar/delivery/internal/usecases/accounts"
	"github.com/indigowar/delivery/pkg/jwt"
)

type Delivery struct {
	finder         accounts.Finder
	registrator    accounts.Registrator
	validator      accounts.CredentialsValidator
	profileUpdater accounts.ProfileUpdater

	tokenValidator *jwt.Validator

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

	d.router.GET("/api/account/id/:id", handlers.FindByIdHandler(d.finder))
	d.router.GET("/api/account/phone/:phone", handlers.FindByPhoneHandler(d.finder))

	d.router.GET("/api/account", handlers.GetOwnInfo(d.finder), middleware.WithJWTAuthentication(d.tokenValidator))
}

func (d *Delivery) AddRegistrator(registrator accounts.Registrator) {
	d.registrator = registrator

	d.router.POST("/api/account", handlers.RegistrationHandler(d.registrator))
}

func (d *Delivery) AddCredentialsValidator(validator accounts.CredentialsValidator) {
	d.validator = validator

	d.router.POST("/api/credentials", handlers.ValidateCredentialsHandler(d.validator))
}

func (d *Delivery) AddProfileUpdater(updater accounts.ProfileUpdater) {
	d.profileUpdater = updater

	// todo: implement profile updater handlers and add them here.
	// d.router.PUT("/api/account", handlers.ProfileUpdateHandler(d.profileUpdater), middleware.WithJWTAuthentication(d.tokenValidator))
}

func NewDelivery(port int, validator *jwt.Validator) *Delivery {
	delivery := &Delivery{
		tokenValidator: validator,
	}

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
