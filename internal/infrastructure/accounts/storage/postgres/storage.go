package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/indigowar/delivery/internal/entities"
	"github.com/indigowar/delivery/internal/usecases/accounts"
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

	db.AutoMigrate(&account{})

	return &Storage{
		db: db,
	}, nil
}

// GetByID - implements Storage
func (s *Storage) GetByID(ctx context.Context, id uuid.UUID) (*entities.Account, error) {
	a, err := s.getByID(id)
	if err != nil {
		return nil, err
	}

	return toEntity(*a), nil
}

// GetByPhone - implements Storage
func (s *Storage) GetByPhone(ctx context.Context, phone string) (*entities.Account, error) {
	a := account{}

	if err := s.db.First(&a, "phone = ?", phone).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, accounts.ErrStorageNotFound
		}

		return nil, fmt.Errorf("failed to find entity: %w", err)
	}

	return toEntity(a), nil
}

// Add - implements Storage
func (s *Storage) Add(ctx context.Context, account *entities.Account) (*entities.Account, error) {
	model := fromEntity(account)

	if err := s.db.Create(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, accounts.ErrStorageAlreadyExists
		}
	}

	a, err := s.getByID(model.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get entity, after creation: %s", err)
	}

	return toEntity(*a), nil
}

// Delete - implements Storage
func (s *Storage) Delete(ctx context.Context, id uuid.UUID) error {
	a := account{}

	if err := s.db.Delete(&a, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return accounts.ErrStorageNotFound
		}
		return fmt.Errorf("failed to delete entity: %w", err)
	}

	return nil
}

// Update - implements Storage
func (s *Storage) Update(ctx context.Context, account *entities.Account) error {
	_, err := s.getByID(account.ID())

	if err == nil {
		return accounts.ErrStorageNotFound
	}

	if !errors.Is(err, accounts.ErrStorageNotFound) {
		return err
	}

	model := fromEntity(account)

	if err := s.db.Save(&model).Error; err != nil {
		return err
	}

	return nil
}

func (s *Storage) getByID(id uuid.UUID) (*account, error) {
	var a account

	if err := s.db.First(&a, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, accounts.ErrStorageNotFound
		}

		return nil, fmt.Errorf("failed to find entity, because: %w", err)
	}

	return &a, nil
}

func (s *Storage) Close() error {
	db, err := s.db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}
