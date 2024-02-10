package rest

import (
	"github.com/labstack/echo/v4"

	"github.com/indigowar/delivery/internal/usecases/accounts"
)

type findByPhoneRequest struct{}
type findByPhoneResponse struct{}

func findByPhoneHandler(f accounts.Finder) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
