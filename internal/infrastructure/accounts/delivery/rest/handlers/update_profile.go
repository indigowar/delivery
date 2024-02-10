package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/indigowar/delivery/internal/usecases/accounts"
)

type profileUpdateRequest struct{}
type profileUpdateResponse struct{}

func ProfileUpdateHandler(svc accounts.ProfileUpdater) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
