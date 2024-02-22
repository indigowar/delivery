package menu

import (
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"

	"github.com/indigowar/delivery/internal/entities"
)

type DishServiceImpl struct {
	logger       *slog.Logger
	dishStorage  DishStorage
	imageStorage ImageStorage
}

var _ DishService = &DishServiceImpl{}

func NewDishService(logger *slog.Logger, storage DishStorage, imgStorage ImageStorage) *DishServiceImpl {
	return &DishServiceImpl{
		logger:       logger,
		dishStorage:  storage,
		imageStorage: imgStorage,
	}
}

func (svc *DishServiceImpl) Get(ctx context.Context, id uuid.UUID) (*entities.Dish, error) {
	dish, err := svc.dishStorage.Get(ctx, id)
	if err != nil {
		if errors.Is(err, ErrDishIsNotInStorage) {
			svc.logger.Info(
				"FAILED RETRIEVE",
				"entity", id.String(),
				"error", err.Error(),
			)

			return nil, ErrDishNotFound
		}

		svc.logger.Warn(
			"FAILED RETRIEVE",
			"entity", id.String(),
			"error", err.Error(),
		)

		return nil, ErrInternalServerError
	}

	return dish, nil
}

func (svc *DishServiceImpl) Create(ctx context.Context, info *DishInfo, image []byte) (*entities.Dish, error) {
	// todo: Implement
	return nil, nil
}

func (svc *DishServiceImpl) Update(ctx context.Context, dish uuid.UUID, info *DishInfo) (*entities.Dish, error) {
	// todo: Implement
	return nil, nil
}

func (svc *DishServiceImpl) UpdateImage(ctx context.Context, dish uuid.UUID, image []byte) (*entities.Dish, error) {
	// todo: Implement
	return nil, nil
}

func (svc *DishServiceImpl) Delete(ctx context.Context, dish uuid.UUID) (*entities.Dish, error) {
	// todo: Implement
	return nil, nil
}
