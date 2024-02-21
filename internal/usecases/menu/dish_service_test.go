package menu

import (
	"bytes"
	"context"
	"errors"
	"log/slog"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	"github.com/indigowar/delivery/internal/entities"
)

type DishServiceTestSuite struct {
	suite.Suite

	loggerBuffer *bytes.Buffer
	logger       *slog.Logger
	dishStorage  *DishStorageMock
	imageStorage *ImageStorageMock

	service *DishServiceImpl
}

func (suite *DishServiceTestSuite) SetupTest() {
	suite.loggerBuffer = &bytes.Buffer{}
	suite.logger = slog.New(slog.NewTextHandler(suite.loggerBuffer, nil))

	suite.dishStorage = &DishStorageMock{}
	suite.imageStorage = &ImageStorageMock{}

	suite.service = NewDishService(suite.logger, suite.dishStorage, suite.imageStorage)
}

func (suite *DishServiceTestSuite) TestGetWhenStorageEmpty() {
	suite.dishStorage.GetFunc = func(ctx context.Context, id uuid.UUID) (*entities.Dish, error) {
		return nil, ErrDishIsNotInStorage
	}

	dish, err := suite.service.Get(context.Background(), uuid.New())

	suite.Nil(dish, "DishService.Get should return nil as entity, when the entity is not found in storage")
	suite.ErrorIsf(err, ErrDishNotFound, "DishService.Get should return ErrDishNotFound as an error, instead of %s", err)
}

func (suite *DishServiceTestSuite) TestGetWhenStorageHasUnexpectedError() {
	suite.dishStorage.GetFunc = func(ctx context.Context, id uuid.UUID) (*entities.Dish, error) {
		return nil, errors.New("unexpected error")
	}

	dish, err := suite.service.Get(context.Background(), uuid.New())

	suite.Nil(dish, "DishService.Get should return nil as entity, when the unexpected error returned from the storage")
	suite.ErrorIsf(err, ErrDishNotFound, "DishService.Get should return ErrInternalServerError as an error, instead of %s", err)
}

func TestDishService(t *testing.T) {
	suite.Run(t, new(DishServiceTestSuite))
}
