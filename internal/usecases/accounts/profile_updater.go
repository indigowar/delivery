package accounts

import (
	"context"
	"errors"
	"log"
	"net/mail"

	"github.com/google/uuid"
	"github.com/indigowar/delivery/internal/entities"
)

// profileUpdater - is an implementation of ProfileUpdater interface
type profileUpdater struct {
	storage      Storage
	imageStorage ImageStorage
}

// LinkEmailToAccount implements ProfileUpdater
func (svc *profileUpdater) LinkEmailToAccount(ctx context.Context, id uuid.UUID, addr *mail.Address) error {
	account, err := svc.storage.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, ErrStorageNotFound) {
			return ErrAccountNotFound
		}

		log.Printf("account storage failed to find user: %s\n", err.Error())
		return ErrInternalServerError
	}

	if err := account.SetEmail(addr); err != nil {
		return ErrProvidedDataIsInvalid
	}

	if err := svc.storage.Update(ctx, account); err != nil {
		log.Printf("account storage failed to update user: %s\n", err.Error())
		return ErrInternalServerError
	}

	return nil
}

// UpdateFirstName implements ProfileUpdater
func (svc *profileUpdater) UpdateFirstName(ctx context.Context, id uuid.UUID, firstName string) error {
	account, err := svc.getAccountByID(ctx, id)
	if err != nil {
		return err
	}

	if err := account.SetFirstName(firstName); err != nil {
		return ErrProvidedDataIsInvalid
	}

	return nil
}

// UpdateSurname implements ProfileUpdater
func (svc *profileUpdater) UpdateSurname(ctx context.Context, id uuid.UUID, surname string) error {
	account, err := svc.getAccountByID(ctx, id)
	if err != nil {
		return err
	}

	if err := account.SetSurname(surname); err != nil {
		return ErrProvidedDataIsInvalid
	}

	return nil
}

// LoadProfileImage implements ProfileUpdater
func (svc *profileUpdater) LoadProfileImage(ctx context.Context, id uuid.UUID, image []byte) error {
	account, err := svc.getAccountByID(ctx, id)
	if err != nil {
		return err
	}

	url, err := svc.imageStorage.Add(ctx, image)
	if err != nil {
		log.Printf("failed to save image in storage:%s\n", err)
		return ErrInternalServerError
	}

	if err := account.SetProfileImageUrl(url); err != nil {
		return ErrProvidedDataIsInvalid
	}

	if err := svc.storage.Update(ctx, account); err != nil {
		log.Printf("failed to update account: %s\n", err)
		return ErrInternalServerError
	}

	return nil
}

func (svc *profileUpdater) getAccountByID(ctx context.Context, id uuid.UUID) (*entities.Account, error) {
	account, err := svc.storage.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, ErrStorageNotFound) {
			return nil, ErrAccountNotFound
		}

		log.Printf("account storage is failed to load user: %s\n", err)
		return nil, ErrInternalServerError
	}

	return account, nil
}

func NewProfileUpdater(storage Storage, imageStorage ImageStorage) ProfileUpdater {
	return &profileUpdater{
		storage:      storage,
		imageStorage: imageStorage,
	}
}
