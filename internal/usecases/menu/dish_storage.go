package menu

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/indigowar/delivery/internal/entities"
)

//go:generate moq -out dish_storage_moq_test.go . DishStorage

var (
	ErrDishIsNotInStorage   = errors.New("dish is not found in storage")
	ErrDishAlreadyInStorage = errors.New("dish is already in the storage")
)

// DishStorage - storage adapter for dish entity.
type DishStorage interface {
	Get(ctx context.Context, id uuid.UUID) (*entities.Dish, error)
	GetSet(ctx context.Context, ids uuid.UUIDs) ([]*entities.Dish, error)
	Add(ctx context.Context, dish *entities.Dish) (*entities.Dish, error)
	Remove(ctx context.Context, id uuid.UUID) (*entities.Dish, error)
	Update(ctx context.Context, dish *entities.Dish) (*entities.Dish, error)
}
