package rest

import (
	"github.com/indigowar/delivery/internal/usecases/accounts"
	"github.com/labstack/echo/v4"
)

type registrationRequest struct{}
type registrationResponse struct{}

func registrationHandler(svc accounts.Registrator) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
