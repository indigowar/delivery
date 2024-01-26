package accounts

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"

	"github.com/indigowar/delivery/internal/entities"
)

// registrator - is an implementation of Registrator interface
type registrator struct {
	storage Storage
}

// RegisterAccount implements Registrator
func (r *registrator) RegisterAccount(ctx context.Context, phone string, password string) (uuid.UUID, error) {
	account, err := entities.NewAccount(phone, password)
	if err != nil {
		return uuid.UUID{}, ErrInvalidCredentials
	}

	account, err = r.storage.Add(ctx, account)
	if err != nil {
		if errors.Is(err, ErrStorageAlreadyExists) {
			return uuid.UUID{}, ErrAccountIsAlreadyExists
		}

		log.Println(err)

		return uuid.UUID{}, err
	}

	return account.ID(), nil
}

func NewRegistrator(storage Storage) Registrator {
	return &registrator{
		storage: storage,
	}
}
