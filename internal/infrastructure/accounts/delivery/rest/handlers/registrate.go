package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/indigowar/delivery/internal/usecases/accounts"
)

type registrationRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type registrationResponse struct {
	ID string `json:"id"`
}

type registrationErrorResponse struct {
	Error string `json:"error"`
	Msg   string `json:"msg"`
}

func RegistrationHandler(svc accounts.Registrator) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := registrationRequest{}
		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, registrationErrorResponse{
				Error: "request data is invalid",
				Msg:   "failed to bind data to request structure",
			})
		}

		id, err := svc.RegisterAccount(c.Request().Context(), request.Phone, request.Password)
		if err != nil {
			if errors.Is(err, accounts.ErrAccountIsAlreadyExists) {
				return c.JSON(http.StatusBadRequest, registrationErrorResponse{
					Error: "account already exists",
					Msg:   "account with provided phone already  exists.",
				})
			}

			if errors.Is(err, accounts.ErrProvidedDataIsInvalid) {
				return c.JSON(http.StatusBadRequest, registrationErrorResponse{
					Error: "data is invalid",
					Msg:   "provided data is invalid",
				})
			}

			return c.JSON(http.StatusInternalServerError, registrationErrorResponse{
				Error: "internal",
				Msg:   "internal server error occurred",
			})
		}

		return c.JSON(http.StatusCreated, registrationResponse{
			ID: id.String(),
		})
	}
}
