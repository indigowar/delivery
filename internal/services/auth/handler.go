package auth

import (
	"errors"
	"log"
	"net/http"

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
	return func(c echo.Context) error {
		phone := c.FormValue("phone")
		password := c.FormValue("password")

		id, err := h.service.Login(c.Request().Context(), phone, password)
		if err != nil {
			if errors.Is(err, ErrCredentialsAreInvalid) {
				return c.NoContent(http.StatusBadRequest)
			}
			return c.NoContent(http.StatusInternalServerError)
		}

		h.sm.Put(c.Request().Context(), "user-id", id.String())
		return c.Redirect(http.StatusSeeOther, "/")
	}
}

func (h *Handler) HandleRegisterRequest() echo.HandlerFunc {
	return func(c echo.Context) error {
		phone := c.FormValue("phone")
		password := c.FormValue("password")

		id, err := h.service.Register(c.Request().Context(), phone, password)
		if err != nil {
			log.Println(err)
			if errors.Is(err, ErrCredentialsAreInvalid) || errors.Is(err, ErrAlreadyInUse) {
				return c.NoContent(http.StatusBadRequest)
			}
			return c.NoContent(http.StatusInternalServerError)
		}

		h.sm.Put(c.Request().Context(), "user-id", id.String())
		return c.Redirect(http.StatusSeeOther, "/")
	}
}
