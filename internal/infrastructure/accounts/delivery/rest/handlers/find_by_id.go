package handlers

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/indigowar/delivery/internal/entities"
	"github.com/indigowar/delivery/internal/usecases/accounts"
)

type findByIdResponse struct {
	ID              string  `json:"id"`
	Phone           string  `json:"phone"`
	Email           *string `json:"email,omitempty"`
	Name            *string `json:"name,omitempty"`
	ProfileImageUrl *string `json:"profile_image,omitempty"`
}

func FindByIdHandler(svc accounts.Finder) echo.HandlerFunc {
	return func(c echo.Context) error {
		paramId := c.Param("id")

		id, err := uuid.Parse(paramId)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		account, err := svc.GetAccount(c.Request().Context(), id)
		if err != nil {
			if errors.Is(err, accounts.ErrAccountNotFound) {
				return c.NoContent(http.StatusNotFound)
			}

			return c.NoContent(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusFound, makeFindByIdResponse(account))
	}
}

func makeFindByIdResponse(account *entities.Account) findByIdResponse {
	response := findByIdResponse{
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
