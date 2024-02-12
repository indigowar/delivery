package rest

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/indigowar/delivery/internal/entities"
	"github.com/indigowar/delivery/internal/usecases/auth"
)

func startSessionHandler(svc auth.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request startSessionRequest
		if err := c.Bind(&request); err != nil {
			return errorInvalidRequestData(c)
		}

		tokens, err := svc.StartSession(c.Request().Context(), request.Phone, request.Phone)
		if err != nil {
			if errors.Is(err, auth.ErrInvalidCredentials) {
				return errorInvalidCredentials(c)
			}

			return errorInternal(c)
		}

		return c.JSON(http.StatusCreated, tokenPairResponse{
			Access:  tokens.ShortLiveToken,
			Session: string(tokens.Session.Token),
		})
	}
}

func endSessionHandler(svc auth.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request sessionTokenRequest
		if err := c.Bind(&request); err != nil {
			return errorInvalidRequestData(c)
		}

		err := svc.EndSession(c.Request().Context(), entities.SessionToken(request.Token))
		if err != nil {
			if errors.Is(err, auth.ErrSessionDoesNotExists) {
				return errorNotFound(c)
			}

			return errorInternal(c)
		}

		return c.NoContent(http.StatusOK)
	}
}

func extendSessionHandler(svc auth.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request sessionTokenRequest
		if err := c.Bind(&request); err != nil {
			return errorInvalidRequestData(c)
		}

		tokens, err := svc.ExtendSession(c.Request().Context(), entities.SessionToken(request.Token))
		if err != nil {
			if errors.Is(err, auth.ErrSessionDoesNotExists) {
				return errorNotFound(c)
			}

			return errorInternal(c)
		}

		return c.JSON(http.StatusOK, tokenPairResponse{
			Access:  tokens.ShortLiveToken,
			Session: string(tokens.Session.Token),
		})
	}
}

func getAccessTokenHandler(svc auth.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request sessionTokenRequest
		if err := c.Bind(&request); err != nil {
			return errorInvalidRequestData(c)
		}

		access, err := svc.GetAccessToken(c.Request().Context(), entities.SessionToken(request.Token))
		if err != nil {
			if errors.Is(err, auth.ErrSessionDoesNotExists) {
				return errorNotFound(c)
			}

			return errorInternal(c)
		}

		return c.JSON(http.StatusOK, accessTokenResponse{
			Access: access,
		})
	}
}

func errorInvalidRequestData(c echo.Context) error {
	return c.JSON(http.StatusBadRequest, errorResponse{
		Error: "bad request",
		Msg:   "failed to bind data to request structure",
	})
}

func errorInternal(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, errorResponse{
		Error: "internal",
		Msg:   "internal server error occurred",
	})
}

func errorInvalidCredentials(c echo.Context) error {
	return c.JSON(http.StatusBadRequest, errorResponse{
		Error: "invalid credentials",
		Msg:   "provided credentials are invalid",
	})
}

func errorNotFound(c echo.Context) error {
	return c.JSON(http.StatusNotFound, errorResponse{
		Error: "not found",
		Msg:   "session with provided token is not found",
	})
}
