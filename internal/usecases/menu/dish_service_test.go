package menu

import (
	"bytes"
	"context"
	"errors"
	"log/slog"
	"net/url"
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
	if suite.NotNil(err, "DishService.Get should return an error") {
		suite.ErrorIsf(err, ErrDishNotFound, "DishService.Get should return ErrDishNotFound as an error, instead of %s", err)
	}
}

func (suite *DishServiceTestSuite) TestGetWhenStorageHasUnexpectedError() {
	suite.dishStorage.GetFunc = func(ctx context.Context, id uuid.UUID) (*entities.Dish, error) {
		return nil, errors.New("unexpected error")
	}

	dish, err := suite.service.Get(context.Background(), uuid.New())

	suite.Nil(dish, "DishService.Get should return nil as entity, when the unexpected error returned from the storage")
	if suite.NotNil(err, "DishService.Get should return an error") {
		suite.ErrorIsf(err, ErrDishNotFound, "DishService.Get should return ErrInternalServerError as an error, instead of %s", err)
	}
}

func (suite *DishServiceTestSuite) TestCreateWhenImageStorageCantSaveImage() {
	suite.imageStorage.AddFunc = func(ctx context.Context, image []byte) (*url.URL, error) {
		return nil, errors.New("failed to save")
	}

	suite.dishStorage.AddFunc = func(ctx context.Context, dish *entities.Dish) (*entities.Dish, error) {
		return dish, nil
	}

	createInfo := DishInfo{
		Name:  "soup",
		Price: 60.0,
		Ingredients: []string{
			"water",
			"carrot",
			"cabbage",
		},
	}

	image := make([]byte, 10)

	dish, err := suite.service.Create(context.Background(), &createInfo, image)

	suite.Nil(dish, "DishService.Create should return nil as entity, when the ImageStorage can't save the image")
	if suite.NotNil(err, "DishService.Create should return an error") {
		suite.ErrorIsf(err, ErrInternalServerError, "DishService.Create should return ErrInternalServerError as an error, instead of %s", err)
	}
}

func (suite *DishServiceTestSuite) TestCreateWhenDishStorageHasUnexpectedError() {
	suite.imageStorage.AddFunc = func(ctx context.Context, image []byte) (*url.URL, error) {
		return url.Parse("http://valid_url.com")
	}

	suite.dishStorage.AddFunc = func(ctx context.Context, dish *entities.Dish) (*entities.Dish, error) {
		return nil, errors.New("unexpected error")
	}

	createInfo := DishInfo{
		Name:  "soup",
		Price: 60.0,
		Ingredients: []string{
			"water",
			"carrot",
			"cabbage",
		},
	}

	image := make([]byte, 10)

	dish, err := suite.service.Create(context.Background(), &createInfo, image)

	suite.Nil(dish, "DishService.Create should return nil as entity, when DishStorage can't save the dish info")
	if suite.NotNil(err, "DishService.Create should return an error") {
		suite.ErrorIsf(err, ErrInternalServerError, "DishService.Create should return ErrInternalServerError as an error, instead of %s", err)
	}
}

func (suite *DishServiceTestSuite) TestCreateValid() {
	validUrl, _ := url.Parse("http://valid_url.com")

	suite.imageStorage.AddFunc = func(ctx context.Context, image []byte) (*url.URL, error) {
		return validUrl, nil
	}

	suite.dishStorage.AddFunc = func(ctx context.Context, dish *entities.Dish) (*entities.Dish, error) {
		return dish, nil
	}

	createInfo := DishInfo{
		Name:  "soup",
		Price: 60.0,
		Ingredients: []string{
			"water",
			"carrot",
			"cabbage",
		},
	}

	image := make([]byte, 60)

	dish, err := suite.service.Create(context.Background(), &createInfo, image)

	suite.Nilf(err, "DishService.Create shouldn't return an error, when input is valid")

	if suite.NotNil(dish, "DishService.Create should not return nil dish, when everything is valid") {
		suite.Equalf(dish.Name(), createInfo.Name, "DishService.Create expected dish name %s, got %s", createInfo.Name, dish.Name())
		suite.Equalf(dish.Price(), createInfo.Price, "DishService.Create expected dish price %f. got %f", createInfo.Price, dish.Price())
		suite.ElementsMatch(dish.Ingredients, createInfo.Ingredients, "DishService.Create ingredients should be the same as in createInfo")
		suite.Equal(dish.Image(), validUrl, "DishService.Create should have url from the ImageStorage")
	}
}

func TestDishService(t *testing.T) {
	suite.Run(t, new(DishServiceTestSuite))
}
