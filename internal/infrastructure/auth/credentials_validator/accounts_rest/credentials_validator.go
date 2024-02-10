package accountsrest

import (
	"context"
	"net/url"

	"github.com/google/uuid"
)

type CredentialsValidator struct {
	host *url.URL
}

func (cv *CredentialsValidator) Validate(ctx context.Context, phone string, password string) (uuid.UUID, error) {
	panic("not implemented")
}

func NewCredentialsValidator(host *url.URL) *CredentialsValidator {
	return &CredentialsValidator{
		host: host,
	}
}
