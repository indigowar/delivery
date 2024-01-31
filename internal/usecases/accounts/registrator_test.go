package accounts

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/indigowar/delivery/internal/entities"
	"github.com/stretchr/testify/suite"
)

type RegistratorTestSuite struct {
	suite.Suite

	storage     *StorageMock
	registrator Registrator
}

func (s *RegistratorTestSuite) SetupTest() {
	s.storage = &StorageMock{}

	s.storage.AddFunc = func(ctx context.Context, account *entities.Account) (*entities.Account, error) {
		return nil, nil
	}

	s.registrator = NewRegistrator(s.storage)
}

func (s *RegistratorTestSuite) TestRegisterAccountWithInvalidData() {
	invalidData := []struct {
		phone    string
		password string
	}{
		{phone: "", password: ""},
		{phone: "addsa", password: "hi"},
		{phone: "12341431", password: "notgoodpassword"},
	}

	for _, data := range invalidData {
		id, err := s.registrator.RegisterAccount(context.TODO(), data.phone, data.password)

		s.Equal(id, uuid.UUID{}, "Registrator.RegisterAccount should return uuid.UUID{}, when failed to create an account")
		s.NotNil(err, "Registrator.RegisterAccount should return an error, when data is invalid")
		s.ErrorIsf(err, ErrInvalidCredentials, "Registrator.RegisterAccount should return ErrInvalidCredentials")
	}
}

func TestRegistrator(t *testing.T) {
	suite.Run(t, new(RegistratorTestSuite))
}
