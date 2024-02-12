package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/indigowar/delivery/internal/usecases/accounts"
)

type validateCredentialsRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type validateCredentialsResponse struct {
	Id string `json:"id"`
}

type validateCredentialsErrorResponse struct {
	Error string `json:"error"`
	Msg   string `json:"msg"`
}

func ValidateCredentialsHandler(v accounts.CredentialsValidator) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := validateCredentialsRequest{}

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, &validateCredentialsErrorResponse{
				Error: "bad request",
				Msg:   "failed to bind request data to request structure",
			})
		}

		id, err := v.ValidateCredentials(c.Request().Context(), request.Phone, request.Password)
		if err != nil {
			if errors.Is(err, accounts.ErrInvalidCredentials) {
				return c.JSON(http.StatusBadRequest, &validateCredentialsErrorResponse{
					Error: "invalid credentials",
					Msg:   "provided credentials are invalid",
				})
			}

			return c.JSON(http.StatusInternalServerError, &validateCredentialsErrorResponse{
				Error: "internal server error",
				Msg:   "an internal server error occurred on service.",
			})
		}

		return c.JSON(http.StatusOK, &validateCredentialsResponse{Id: id.String()})
	}
}
