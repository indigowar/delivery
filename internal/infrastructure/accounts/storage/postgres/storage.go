package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/indigowar/delivery/internal/entities"
	"github.com/indigowar/delivery/internal/usecases/accounts"
	"github.com/indigowar/delivery/pkg/postgres"
)

type storage struct {
	con *pgx.Conn
}

// GetByID - implements Storage
func (s *storage) GetByID(ctx context.Context, id uuid.UUID) (*entities.Account, error) {
	panic("unimplemented")
}

// GetByPhone - implements Storage
func (s *storage) GetByPhone(ctx context.Context, phone string) (*entities.Account, error) {
	panic("unimplemented")
}

// Add - implements Storage
func (s *storage) Add(ctx context.Context, account *entities.Account) (*entities.Account, error) {
	panic("unimplemented")
}

// Delete - implements Storage
func (s *storage) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// Update - implements Storage
func (s *storage) Update(ctx context.Context, account *entities.Account) error {
	panic("unimplemented")
}

func NewStorage(host string, port string, user string, password string, db string) (accounts.Storage, error) {
	con, err := postgres.Connect(host, port, user, password, db)
	if err != nil {
		return nil, err
	}

	return &storage{
		con: con,
	}, nil
}
