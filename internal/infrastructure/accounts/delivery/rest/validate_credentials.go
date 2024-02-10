package rest

import (
	"github.com/labstack/echo/v4"

	"github.com/indigowar/delivery/internal/usecases/accounts"
)

type validateCredentialsRequest struct{}
type validateCredentialsResponse struct{}

func validateCredentialsHandler(v accounts.CredentialsValidator) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
