package middleware

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
)

func LoggedInMiddleware(sm *scs.SessionManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			id := sm.GetString(c.Request().Context(), "user-id")
			if id == "" {
				return c.Redirect(http.StatusSeeOther, "/login")
			}
			return next(c)
		}
	}
}
