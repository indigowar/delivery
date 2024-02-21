package menu

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/indigowar/delivery/internal/entities"
)

//go:generate moq -out category_storage_moq_test.go . CategoryStorage

var (
	ErrCategoryNotFoundInStorage = errors.New("category is not found in the storage")
	ErrCategoryAlreadyExists     = errors.New("category already exists")
)

type CategoryStorage interface {
	Get(ctx context.Context, id uuid.UUID) (*entities.Category, error)
	GetSet(ctx context.Context, ids uuid.UUIDs) ([]*entities.Category, error)
	GetByRestaurant(ctx context.Context, id uuid.UUID) ([]*entities.Category, error)
	Add(ctx context.Context, category *entities.Category) (*entities.Category, error)
	Remove(ctx context.Context, id uuid.UUID) (*entities.Category, error)
}
