package accountsrest

import (
	"context"
	"net/url"
)

type CredentialsValidator struct {
	host *url.URL
}

func (cv *CredentialsValidator) Validate(ctx context.Context, phone string, password string) error {
	panic("not implemented")
}

func NewCredentialsValidator(host *url.URL) *CredentialsValidator {
	return &CredentialsValidator{
		host: host,
	}
}
