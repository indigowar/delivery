package rest

import (
	"github.com/indigowar/delivery/internal/usecases/accounts"
	"github.com/labstack/echo/v4"
)

type findByIdRequest struct{}
type findByIdResponse struct{}

func findByIdHandler(svc accounts.Finder) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
