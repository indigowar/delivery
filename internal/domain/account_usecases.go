package domain

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	// ErrAccountNotFound - returns when the account is not found in the storage
	ErrAccountAlreadyExists = errors.New("account already exists")
	// ErrAccountAlreadyExists - returns when the account already exists in the storage,
	// and the action tries to add it again.
	ErrAccountNotFound = errors.New("account is not found")
)

type GetAccountByIDUseCase interface {
	GetByID(ctx context.Context, id uuid.UUID) (*Account, error)
}

type GetAccountByPhoneUseCase interface {
	GetByPhone(ctx context.Context, phone string) (*Account, error)
}

type CreateAccountUseCase interface {
	Create(ctx context.Context, account *Account) error
}

type UpdateAccountUseCase interface {
	Update(ctx context.Context, account *Account) error
}
