package status

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type response struct {
	Service   string `json:"service"`
	IsRunning bool   `json:"is_running"`
}

func StatusHandler(service string) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, response{
			Service:   service,
			IsRunning: true,
		})
	}
}
