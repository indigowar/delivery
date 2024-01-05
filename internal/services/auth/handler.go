package auth

import (
	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service Service

	sm *scs.SessionManager
}

func NewHandler(service Service, sm *scs.SessionManager) Handler {
	return Handler{
		service: service,
		sm:      sm,
	}
}

func (h *Handler) ServeLoginPage(handle string) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
		return loginPage(handle).Render(c.Request().Context(), c.Response().Writer)
	}
}

func (h *Handler) ServeRegistrationPage(handle string) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
		return registerPage(handle).Render(c.Request().Context(), c.Response().Writer)
	}
}

func (h *Handler) HandleLoginRequest() echo.HandlerFunc {
	return func(c echo.Context) error { return nil }
}

func (h *Handler) HandleRegisterRequest() echo.HandlerFunc {
	return func(c echo.Context) error { return nil }
}
