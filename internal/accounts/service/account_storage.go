package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/indigowar/delivery/internal/accounts/domain"
)

type StorageErrorType = uint

const (
	StorageErrorTypeNotFound StorageErrorType = iota
	StorageErrorTypeAlreadyExists
	StorageErrroTypeOther
)

// AccountStorage - interface to access the accounts in storage
type AccountStorage interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.Account, error)
	GetByPhone(ctx context.Context, phone string) (*domain.Account, error)
	Add(ctx context.Context, account domain.Account) (*domain.Account, error)
	Update(ctx context.Context, account domain.Account) (*domain.Account, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type AccountStorageError struct {
	ty  StorageErrorType
	msg string
}

func (e *AccountStorageError) Error() string {
	switch e.ty {
	case StorageErrorTypeNotFound:
		return fmt.Sprintf("Account is not found %s", e.msg)
	case StorageErrorTypeAlreadyExists:
		return fmt.Sprintf("Account already exists %s", e.msg)
	default:
		return fmt.Sprintf("Unexpected error occurred %s", e.msg)
	}
}

func NewAccountStorageError(ty StorageErrorType, msg string) error {
	return &AccountStorageError{
		ty:  ty,
		msg: msg,
	}
}
