package accountsrest

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type CredentialsValidator struct {
	host *url.URL
}

type request struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type validResponse struct {
	Id string `json:"id"`
}

type errorResponse struct {
	Error string `json:"error"`
	Msg   string `json:"msg"`
}

func (cv *CredentialsValidator) Validate(ctx context.Context, phone string, password string) (uuid.UUID, error) {
	request, err := json.Marshal(request{
		Phone:    phone,
		Password: password,
	})
	if err != nil {
		return uuid.UUID{}, err
	}

	addr := cv.host.String() + "/api/credentials"

	response, err := http.Post(addr, "application/json", bytes.NewBuffer(request))
	if err != nil {
		return uuid.UUID{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		var body errorResponse
		if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
			return uuid.UUID{}, err
		}

		return uuid.UUID{}, errors.New(body.Error)
	}

	var body validResponse
	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		return uuid.UUID{}, err
	}

	id, err := uuid.Parse(body.Id)
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func NewCredentialsValidator(host *url.URL) *CredentialsValidator {
	return &CredentialsValidator{
		host: host,
	}
}
