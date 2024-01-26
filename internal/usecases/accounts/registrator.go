package accounts

import (
	"context"

	"github.com/google/uuid"
)

// registrator - is an implementation of Registrator interface
type registrator struct {
	storage Storage
}

// RegisterAccount implements Registrator
func (r *registrator) RegisterAccount(ctx context.Context, phone string, password string) (uuid.UUID, error) {
	// todo: implement
	panic("unimplemented")
}

func NewRegistrator(storage Storage) Registrator {
	return &registrator{
		storage: storage,
	}
}
