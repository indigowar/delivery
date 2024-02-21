package menu

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/indigowar/delivery/internal/entities"
)

var (
	ErrDishNotFound       = errors.New("dish is not found")
	ErrRestaurantNotFound = errors.New("restaurant is not found")
	ErrDishAlreadyExists  = errors.New("dish already exists")
	ErrInvalidManager     = errors.New("invalid manager")
	ErrInvalidOrderSet    = errors.New("invalid order set was provided")

	ErrProvidedDataIsInvalid = errors.New("provided data is invalid")
	ErrInternalServerError   = errors.New("internal server error")
)

// DishInfo contains all writable fields for entity.Dish
type DishInfo struct {
	Name        string
	About       string
	Ingredients []string
	Price       float64
}

// DishService manages entity.Dish, provides CRUD operations on this entity.
type DishService interface {
	Get(ctx context.Context, id uuid.UUID) (*entities.Dish, error)
	Create(ctx context.Context, info *DishInfo, image []byte) (*entities.Dish, error)
	Update(ctx context.Context, dish uuid.UUID, info *DishInfo) (*entities.Dish, error)
	UpdateImage(ctx context.Context, dish uuid.UUID, image []byte) (*entities.Dish, error)
	Delete(ctx context.Context, dish uuid.UUID) (*entities.Dish, error)
}

// OrderDishSetVerifier verifies that the set of dishes belongs to restaurant
type OrderDishSetVerifier interface {
	Verify(ctx context.Context, dishes []uuid.UUID, restaurant uuid.UUID) error
}

// RestaurantService - manages info about restaurants in the menu service
type RestaurantService interface {
	GetCategories(ctx context.Context, restaurant uuid.UUID) ([]*entities.Category, error)
	GetDishes(ctx context.Context, restaurant uuid.UUID) ([]*entities.Dish, error)

	GetDishesForCategory(ctx context.Context, restaurant uuid.UUID, category uuid.UUID) ([]*entities.Dish, error)
	AddRestaurant(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
}

// CategoryService - manages info about categories.
type CategoryService interface {
	Get(ctx context.Context, id uuid.UUID) (*entities.Category, error)
	GetForRestaurant(ctx context.Context, id uuid.UUID) ([]uuid.UUID, error)
	Create(ctx context.Context, restaurant uuid.UUID, name string, image []byte) (*entities.Category, error)
	UpdateName(ctx context.Context, category uuid.UUID, name string) (*entities.Category, error)
	UpdateImage(ctx context.Context, category uuid.UUID, image []byte) (*entities.Category, error)
	AddDishToCategory(ctx context.Context, category uuid.UUID, dish uuid.UUID) error
	RemoveDishFromCategory(ctx context.Context, category uuid.UUID, dish uuid.UUID) error
}
