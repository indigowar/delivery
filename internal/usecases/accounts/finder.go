package accounts

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"

	"github.com/indigowar/delivery/internal/entities"
)

// finder - is an implementation of Finder interface.
type finder struct {
	storage Storage
}

// GetAccount implements Finder
func (f *finder) GetAccount(ctx context.Context, id uuid.UUID) (*entities.Account, error) {
	account, err := f.storage.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, ErrStorageNotFound) {
			return nil, ErrAccountNotFound
		}

		log.Println(err)

		return nil, ErrInternalServerError
	}

	return account, nil
}

// GetAccountByPhone implements  Finder
func (f *finder) GetAccountByPhone(ctx context.Context, phone string) (*entities.Account, error) {
	if entities.ValidatePhoneNumberValue(phone) != nil {
		return nil, ErrProvidedDataIsInvalid
	}

	account, err := f.storage.GetByPhone(ctx, phone)
	if err != nil {
		if errors.Is(err, ErrStorageNotFound) {
			return nil, ErrAccountNotFound
		}

		log.Println(err)

		return nil, ErrInternalServerError
	}

	return account, nil
}

func NewFinder(storage Storage) Finder {
	return &finder{
		storage: storage,
	}
}
