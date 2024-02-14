package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/indigowar/delivery/internal/entities"
	"github.com/indigowar/delivery/internal/usecases/accounts"
)

type findByPhoneResponse struct {
	ID              string  `json:"id"`
	Phone           string  `json:"phone"`
	Email           *string `json:"email,omitempty"`
	Name            *string `json:"name,omitempty"`
	ProfileImageUrl *string `json:"profile_image,omitempty"`
}

func FindByPhoneHandler(f accounts.Finder) echo.HandlerFunc {
	return func(c echo.Context) error {
		phone := c.Param("phone")

		account, err := f.GetAccountByPhone(c.Request().Context(), phone)
		if err != nil {
			if errors.Is(err, accounts.ErrAccountNotFound) {
				return c.NoContent(http.StatusNotFound)
			}

			return c.NoContent(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusFound, makeFindByPhoneResponse(account))
	}
}

func makeFindByPhoneResponse(account *entities.Account) findByPhoneResponse {
	response := findByPhoneResponse{
		ID:    account.ID().String(),
		Phone: account.Phone(),
	}

	if account.HasEmail() {
		*response.Email = account.Email().String()
	}

	if account.HasName() {
		*response.Name = account.Name()
	}

	if account.HasProfileImageUrl() {
		*response.ProfileImageUrl = account.ProfileImageUrl().String()
	}

	return response
}
