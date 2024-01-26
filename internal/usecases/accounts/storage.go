package accounts

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/indigowar/delivery/internal/entities"
)

//go:generate moq -out storage_moq_test.go . Storage

var (
	ErrStorageNotFound      = errors.New("not found in the storage")
	ErrStorageAlreadyExists = errors.New("already exists in storage")
)

type Storage interface {
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Account, error)
	GetByPhone(ctx context.Context, phone string) (*entities.Account, error)
	Add(ctx context.Context, account *entities.Account) (*entities.Account, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, account *entities.Account) error
}
