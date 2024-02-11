package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/indigowar/delivery/pkg/jwt"
)

// WithJWTAuthentication is a simple middleware that validates JWT Token,
// if it's valid it adds it into request context, otherwise returns an error.
func WithJWTAuthentication(validator *jwt.Validator) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Authorization")
			if token == "" {
				return c.NoContent(http.StatusUnauthorized)
			}

			id, err := validator.Validate(token)
			if err != nil {
				return c.NoContent(http.StatusUnauthorized)
			}

			c.Set("id", id)
			return next(c)
		}
	}
}
