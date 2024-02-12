package handlers

import (
	"errors"
	"net/http"
	"net/mail"
	"net/url"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/indigowar/delivery/internal/usecases/accounts"
)

type findByIdResponse struct {
	ID              uuid.UUID     `json:"id"`
	Phone           string        `json:"phone"`
	Email           *mail.Address `json:"email,omitempty"`
	FirstName       *string       `json:"first_name,omitempty"`
	Surname         *string       `json:"surname,omitempty"`
	ProfileImageUrl *url.URL      `json:"profile_image,omitempty"`
}

type findByIDErrorResponse struct {
	Error string `json:"error"`
	Msg   string `json:"msg"`
}

func FindByIdHandler(svc accounts.Finder) echo.HandlerFunc {
	return func(c echo.Context) error {
		paramId := c.Param("id")

		id, err := uuid.Parse(paramId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, findByIDErrorResponse{
				Error: "bad request",
				Msg:   "given id in params is invalid",
			})
		}

		account, err := svc.GetAccount(c.Request().Context(), id)
		if err != nil {
			if errors.Is(err, accounts.ErrAccountNotFound) {
				return c.JSON(http.StatusNotFound, findByIDErrorResponse{
					Error: "not found",
					Msg:   "accoun with given id is not found",
				})
			}

			return c.NoContent(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, findByIdResponse{
			ID:              account.ID(),
			Phone:           account.Phone(),
			Email:           account.Email(),
			FirstName:       account.FirstName(),
			Surname:         account.Surname(),
			ProfileImageUrl: account.ProfileImageUrl(),
		})
	}
}
