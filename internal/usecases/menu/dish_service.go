package menu

import (
	"context"
	"errors"
	"log/slog"
	"net/url"

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
			svc.logRetrieveFailed(slog.LevelInfo, id, err.Error())
			return nil, ErrDishNotFound
		}

		svc.logRetrieveFailed(slog.LevelWarn, id, err.Error())
		return nil, ErrInternalServerError
	}

	return dish, nil
}

func (svc *DishServiceImpl) Create(ctx context.Context, info *DishInfo, image []byte) (*entities.Dish, error) {
	if err := svc.validateCreateInfo(info, image); err != nil {
		return nil, err
	}

	dish, err := entities.NewDish(*info.Name, *info.Price, nil)
	if err != nil {
		svc.logDishCreateFailed("invalid data provided", err.Error())
		return nil, ErrProvidedDataIsInvalid
	}

	if info.Ingredients != nil {
		dish.Ingredients = *info.Ingredients
	}

	if info.About != nil {
		dish.About = *info.About
	}

	imageUrl, err := svc.uploadImage(ctx, image)
	if err != nil {
		return nil, err
	}

	if err := dish.SetImage(imageUrl); err != nil {
		svc.logDishCreateFailed("image url is invalid", "url is invalid after uploading the image")
		return nil, ErrInternalServerError
	}

	dish, err = svc.dishStorage.Add(ctx, dish)
	if err != nil {
		svc.logDishCreateFailed("failed to add in storage", err.Error())
		return nil, ErrInternalServerError
	}

	return dish, nil
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

func (svc *DishServiceImpl) validateCreateInfo(info *DishInfo, image []byte) error {
	if info.Name == nil || info.Price == nil || image == nil {
		svc.logDishCreateFailed(ErrProvidedDataIsInvalid.Error(), "dish info contains undefined fields")
		return ErrProvidedDataIsInvalid
	}

	if len(image) == 0 {
		svc.logDishCreateFailed(ErrProvidedDataIsInvalid.Error(), "image size is 0")
		return ErrProvidedDataIsInvalid
	}

	return nil
}

func (svc *DishServiceImpl) uploadImage(ctx context.Context, image []byte) (*url.URL, error) {
	url, err := svc.imageStorage.Add(ctx, image)
	if err != nil {
		svc.logDishCreateFailed("failed to upload an image", err.Error())
		return nil, ErrInternalServerError
	}

	return url, nil
}

func (svc *DishServiceImpl) logDishCreateFailed(err string, more string) {
	svc.logger.Info(
		"DISH_CREATE_FAILED",
		"error", err,
		"more", more,
	)
}

func (svc *DishServiceImpl) logRetrieveFailed(level slog.Level, entity uuid.UUID, err string) {
	svc.logger.Log(
		context.Background(),
		level,
		"RETRIEVE_FAILED",

		"entity", entity.ID(),
		"error", err,
	)
}
