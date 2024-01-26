package accounts

import (
	"context"
	"errors"
)

type credentialsValidator struct {
	storage Storage
}

// ValidateCredentials - implements CredentialsValidator
func (c *credentialsValidator) ValidateCredentials(ctx context.Context, phone string, password string) error {
	account, err := c.storage.GetByPhone(ctx, phone)
	if err != nil {
		if errors.Is(err, ErrAccountNotFound) {
			return ErrInvalidCredentials
		}
		return ErrInternalServerError
	}

	if !account.ValidatePassword(password) {
		return ErrInvalidCredentials
	}

	return nil
}

func NewCredentialsValidator(storage Storage) CredentialsValidator {
	return &credentialsValidator{
		storage: storage,
	}
}
