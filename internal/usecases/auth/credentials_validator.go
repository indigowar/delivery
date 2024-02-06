package auth

import "context"

//go:generate moq -out credentials_validator_moq_test.go . CredentialsValidator

type CredentialsValidator interface {
	Validate(ctx context.Context, phone string, password string) error
}
