package accounts

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type credentialsValidator struct {
	storage Storage
}

// ValidateCredentials - implements CredentialsValidator
func (c *credentialsValidator) ValidateCredentials(ctx context.Context, phone string, password string) (uuid.UUID, error) {
	account, err := c.storage.GetByPhone(ctx, phone)
	if err != nil {
		if errors.Is(err, ErrAccountNotFound) {
			return uuid.UUID{}, ErrInvalidCredentials
		}
		return uuid.UUID{}, ErrInternalServerError
	}

	if !account.ValidatePassword(password) {
		return uuid.UUID{}, ErrInvalidCredentials
	}

	return account.ID(), nil
}

func NewCredentialsValidator(storage Storage) CredentialsValidator {
	return &credentialsValidator{
		storage: storage,
	}
}
