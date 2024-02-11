package handlers

import (
	"net/http"
	"net/mail"
	"net/url"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/indigowar/delivery/internal/entities"
	"github.com/indigowar/delivery/internal/usecases/accounts"
)

type ownInfoResponse struct {
	ID              uuid.UUID     `json:"id"`
	Phone           string        `json:"phone"`
	Email           *mail.Address `json:"email,omitempty"`
	FirstName       *string       `json:"first_name,omitempty"`
	Surname         *string       `json:"surname,omitempty"`
	ProfileImageUrl *url.URL      `json:"profile_image,omitempty"`
}

func GetOwnInfo(f accounts.Finder) echo.HandlerFunc {
	fromModel := func(account *entities.Account) ownInfoResponse {
		return ownInfoResponse{
			ID:              account.ID(),
			Phone:           account.Phone(),
			Email:           account.Email(),
			FirstName:       account.FirstName(),
			Surname:         account.Surname(),
			ProfileImageUrl: account.ProfileImageUrl(),
		}
	}

	return func(c echo.Context) error {
		id, ok := c.Get("id").(uuid.UUID)
		if !ok {
			return c.NoContent(http.StatusBadRequest)
		}

		account, err := f.GetAccount(c.Request().Context(), id)
		if err != nil {
			return c.NoContent(http.StatusNotFound)
		}

		return c.JSON(http.StatusOK, fromModel(account))
	}
}
