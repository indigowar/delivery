package postgres

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/indigowar/delivery/internal/entities"
	"github.com/indigowar/delivery/pkg/postgres"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage(host, port, user, password, dbName string) (*Storage, error) {
	db, err := postgres.Connect(host, port, user, password, dbName)
	if err != nil {
		return nil, err
	}

	return &Storage{
		db: db,
	}, nil
}

// GetByID - implements Storage
func (s *Storage) GetByID(ctx context.Context, id uuid.UUID) (*entities.Account, error) {
	panic("unimplemented")
}

// GetByPhone - implements Storage
func (s *Storage) GetByPhone(ctx context.Context, phone string) (*entities.Account, error) {
	panic("unimplemented")
}

// Add - implements Storage
func (s *Storage) Add(ctx context.Context, account *entities.Account) (*entities.Account, error) {
	panic("unimplemented")
}

// Delete - implements Storage
func (s *Storage) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// Update - implements Storage
func (s *Storage) Update(ctx context.Context, account *entities.Account) error {
	panic("unimplemented")
}

func (s *Storage) Close() error {
	db, err := s.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
