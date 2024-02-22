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

func TestDishService(t *testing.T) {
	suite.Run(t, new(DishServiceTestSuite))
}

type DishServiceTestSuite struct {
	suite.Suite

	loggerBuffer *bytes.Buffer
	logger       *slog.Logger
	dishStorage  *DishStorageMock
	imageStorage *ImageStorageMock

	service *DishServiceImpl

	validUrl *url.URL

	getNotFound   func(_ context.Context, id uuid.UUID) (*entities.Dish, error)
	getUnexpected func(_ context.Context, id uuid.UUID) (*entities.Dish, error)
	getValid      func(_ context.Context, id uuid.UUID) (*entities.Dish, error)

	addUnexpected func(_ context.Context, dish *entities.Dish) (*entities.Dish, error)
	addValid      func(_ context.Context, dish *entities.Dish) (*entities.Dish, error)

	addImageError func(_ context.Context, image []byte) (*url.URL, error)
	addImageValid func(value *url.URL) func(_ context.Context, image []byte) (*url.URL, error)

	updateNotFound func(_ context.Context, dish *entities.Dish) (*entities.Dish, error)
	updateValid    func(_ context.Context, dish *entities.Dish) (*entities.Dish, error)
}

func (suite *DishServiceTestSuite) SetupTest() {
	suite.loggerBuffer = &bytes.Buffer{}
	suite.logger = slog.New(slog.NewTextHandler(suite.loggerBuffer, nil))

	suite.dishStorage = &DishStorageMock{}
	suite.imageStorage = &ImageStorageMock{}

	suite.service = NewDishService(suite.logger, suite.dishStorage, suite.imageStorage)

	suite.setupGetFunctions()
	suite.setupAddFunctions()
	suite.setupImageStorageFunctions()
	suite.setupUpdateFunctions()
}

func (suite *DishServiceTestSuite) TestGetWhenStorageEmpty() {
	suite.dishStorage.GetFunc = suite.getNotFound

	dish, err := suite.service.Get(context.Background(), uuid.New())

	suite.Nil(dish, "DishService.Get should return nil as entity, when the entity is not found in storage")
	if suite.NotNil(err, "DishService.Get should return an error") {
		suite.ErrorIsf(err, ErrDishNotFound, "DishService.Get should return ErrDishNotFound as an error, instead of %s", err)
	}
}

func (suite *DishServiceTestSuite) TestGetWhenStorageHasUnexpectedError() {
	suite.dishStorage.GetFunc = suite.getUnexpected

	dish, err := suite.service.Get(context.Background(), uuid.New())

	suite.Nil(dish, "DishService.Get should return nil as entity, when the unexpected error returned from the storage")
	if suite.NotNil(err, "DishService.Get should return an error") {
		suite.ErrorIsf(err, ErrDishNotFound, "DishService.Get should return ErrInternalServerError as an error, instead of %s", err)
	}
}

func (suite *DishServiceTestSuite) TestCreateWhenImageStorageCantSaveImage() {
	suite.imageStorage.AddFunc = suite.addImageError
	suite.dishStorage.AddFunc = suite.addValid

	name := "soup"
	price := 60.0
	ingredients := []string{
		"water",
		"carrot",
		"cabbage",
	}

	createInfo := DishInfo{
		Name:        &name,
		Price:       &price,
		Ingredients: &ingredients,
	}

	image := make([]byte, 10)

	dish, err := suite.service.Create(context.Background(), &createInfo, image)

	suite.Nil(dish, "DishService.Create should return nil as entity, when the ImageStorage can't save the image")
	if suite.NotNil(err, "DishService.Create should return an error") {
		suite.ErrorIsf(err, ErrInternalServerError, "DishService.Create should return ErrInternalServerError as an error, instead of %s", err)
	}
}

func (suite *DishServiceTestSuite) TestCreateWhenDishStorageHasUnexpectedError() {
	validUrl, _ := url.Parse("https://valid_url.com")
	suite.imageStorage.AddFunc = suite.addImageValid(validUrl)
	suite.dishStorage.AddFunc = suite.addUnexpected

	name := "soup"
	price := 60.0
	ingredients := []string{
		"water",
		"carrot",
		"cabbage",
	}

	createInfo := DishInfo{
		Name:        &name,
		Price:       &price,
		Ingredients: &ingredients,
	}

	image := make([]byte, 10)

	dish, err := suite.service.Create(context.Background(), &createInfo, image)

	suite.Nil(dish, "DishService.Create should return nil as entity, when DishStorage can't save the dish info")
	if suite.NotNil(err, "DishService.Create should return an error") {
		suite.ErrorIsf(err, ErrInternalServerError, "DishService.Create should return ErrInternalServerError as an error, instead of %s", err)
	}
}

func (suite *DishServiceTestSuite) TestCreateValid() {
	validUrl, _ := url.Parse("https://valid_url.com")
	suite.imageStorage.AddFunc = suite.addImageValid(validUrl)
	suite.dishStorage.AddFunc = suite.addValid

	name := "soup"
	price := 60.0
	ingredients := []string{
		"water",
		"carrot",
		"cabbage",
	}

	createInfo := DishInfo{
		Name:        &name,
		Price:       &price,
		Ingredients: &ingredients,
	}

	image := make([]byte, 60)

	dish, err := suite.service.Create(context.Background(), &createInfo, image)

	suite.Nilf(err, "DishService.Create shouldn't return an error, when input is valid")

	if suite.NotNil(dish, "DishService.Create should not return nil dish, when everything is valid") {
		suite.Equal(dish.Name(), name, "DishService.Create unexpected field value")
		suite.Equal(dish.Price(), price, "DishService.Create unexpected field value")
		suite.Equal(dish.Image(), validUrl, "DishService.Create unexpected field value")
		suite.ElementsMatch(dish.Ingredients, ingredients, "DishService.Create ingredients should be the same as in createInfo")
	}
}

func (suite *DishServiceTestSuite) TestCreateWithEmptyDishInfo() {
	validUrl, _ := url.Parse("https://valid_url.com")

	suite.imageStorage.AddFunc = suite.addImageValid(validUrl)
	suite.dishStorage.AddFunc = suite.addValid

	createInfo := DishInfo{}

	image := make([]byte, 60)

	dish, err := suite.service.Create(context.Background(), &createInfo, image)

	if suite.NotNil(err, "DishService.Create should return an error, when DishInfo is incomplete") {
		suite.ErrorIs(err, ErrProvidedDataIsInvalid, "DishService.Create should return ErrProvidedDataIsInvalid when DishInfo is incomplete")
	}

	suite.Nil(dish, "DishService.Create should return nil as entity, if DishInfo is invalid")
}

func (suite *DishServiceTestSuite) TestCreateWithDishInfoWithoutName() {
	validUrl, _ := url.Parse("https://valid_url.com")

	suite.imageStorage.AddFunc = suite.addImageValid(validUrl)
	suite.dishStorage.AddFunc = suite.addValid

	createInfo := DishInfo{}

	image := make([]byte, 60)

	dish, err := suite.service.Create(context.Background(), &createInfo, image)

	if suite.NotNil(err, "DishService.Create should return an error, when DishInfo is incomplete") {
		suite.ErrorIs(err, ErrProvidedDataIsInvalid, "DishService.Create should return ErrProvidedDataIsInvalid when DishInfo is incomplete")
	}

	suite.Nil(dish, "DishService.Create should return nil as entity, if DishInfo is invalid")
}

func (suite *DishServiceTestSuite) TestUpdateNotFoundInStorage() {
	suite.dishStorage.GetFunc = suite.getNotFound
	suite.dishStorage.UpdateFunc = suite.updateNotFound

	id := uuid.New()
	name := "hello, world"
	updateInfo := DishInfo{
		Name: &name,
	}

	dish, err := suite.service.Update(context.Background(), id, &updateInfo)

	suite.Nil(dish, "DishService.Update should return nil as entity, when entity is not found")
	if suite.NotNil(err, "DishService.Update should return an error, when the entity is not found") {
		suite.ErrorIs(err, ErrDishNotFound, "DishService.Update should return ErrDishIsNotFound, when the entity is not found")
	}
}

func (suite *DishServiceTestSuite) TestUpdateNotFoundOnUpdateInStorage() {
	suite.dishStorage.GetFunc = func(_ context.Context, id uuid.UUID) (*entities.Dish, error) {
		validUrl, _ := url.Parse("https://valid_url.com")
		d, _ := entities.NewDish("soup", 60.0, validUrl)
		return d, nil
	}

	suite.dishStorage.UpdateFunc = suite.updateNotFound

	id := uuid.New()
	name := "hello, world"
	updateInfo := DishInfo{
		Name: &name,
	}

	dish, err := suite.service.Update(context.Background(), id, &updateInfo)

	suite.Nil(dish, "DishService.Update should return nil as entity, when entity is not found")
	if suite.NotNil(err, "DishService.Update should return an error, when the entity is not found") {
		suite.ErrorIs(err, ErrDishNotFound, "DishService.Update should return ErrDishIsNotFound, when the entity is not found")
	}
}

func (suite *DishServiceTestSuite) TestUpdateWithEmptyInfo() {
	suite.dishStorage.UpdateFunc = suite.updateValid

	id := uuid.New()
	updateInfo := DishInfo{}

	dish, err := suite.service.Update(context.Background(), id, &updateInfo)

	suite.Nil(dish, "DishService.Update should return nil as entity, when entity is not found")
	if suite.NotNil(err, "DishService.Update should return an error, when the entity is not found") {
		suite.ErrorIs(err, ErrProvidedDataIsInvalid, "DishService.Update should return ErrDishIsNotFound, when the DishInfo is empty")
	}
}

func (suite *DishServiceTestSuite) setupAddFunctions() {
	suite.addValid = func(_ context.Context, dish *entities.Dish) (*entities.Dish, error) {
		return dish, nil
	}

	suite.addUnexpected = func(_ context.Context, dish *entities.Dish) (*entities.Dish, error) {
		return nil, errors.New("unexpected error")
	}
}

func (suite *DishServiceTestSuite) setupGetFunctions() {
	suite.getNotFound = func(_ context.Context, _ uuid.UUID) (*entities.Dish, error) {
		return nil, ErrDishIsNotInStorage
	}

	suite.getValid = func(_ context.Context, id uuid.UUID) (*entities.Dish, error) {
		validUrl, _ := url.Parse("https://valid_url.com")
		d, _ := entities.NewDish("soup", 60.0, validUrl)
		return d, nil
	}
}

func (suite *DishServiceTestSuite) setupImageStorageFunctions() {
	suite.addImageError = func(_ context.Context, _ []byte) (*url.URL, error) {
		return nil, errors.New("failed to save")
	}

	suite.addImageValid = func(v *url.URL) func(_ context.Context, _ []byte) (*url.URL, error) {
		return func(_ context.Context, image []byte) (*url.URL, error) {
			return v, nil
		}
	}
}

func (suite *DishServiceTestSuite) setupUpdateFunctions() {
	suite.updateNotFound = func(_ context.Context, _ *entities.Dish) (*entities.Dish, error) {
		return nil, ErrDishIsNotInStorage
	}

	suite.updateValid = func(_ context.Context, dish *entities.Dish) (*entities.Dish, error) {
		return dish, nil
	}
}
