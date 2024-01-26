package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/indigowar/delivery/internal/entities"
	"github.com/indigowar/delivery/internal/usecases/accounts"
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

func NewStorage(host string, port string, user string, password string, db string) accounts.Storage {
	panic("unimplemented")
	// todo: implement

	// - make a connection to db
	// - execute migrations
	// - return a storage implementation
}

func newStorage(con *pgx.Conn) accounts.Storage {
	return &storage{
		con: con,
	}
}
