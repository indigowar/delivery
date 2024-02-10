package auth

import (
	"context"

	"github.com/google/uuid"
)

//go:generate moq -out credentials_validator_moq_test.go . CredentialsValidator

type CredentialsValidator interface {
	Validate(ctx context.Context, phone string, password string) (uuid.UUID, error)
}
