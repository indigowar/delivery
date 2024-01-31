package accounts

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	"github.com/indigowar/delivery/internal/entities"
)

type FinderTestSuite struct {
	suite.Suite

	finder  Finder
	storage *StorageMock
}

func (s *FinderTestSuite) SetupTest() {
	s.storage = &StorageMock{}
	s.finder = NewFinder(s.storage)
}

func (s *FinderTestSuite) TestGetAccountWhenAccountDoesNotExists() {
	s.storage.GetByIDFunc = func(ctx context.Context, id uuid.UUID) (*entities.Account, error) {
		return nil, ErrStorageNotFound
	}

	id := uuid.New()

	account, err := s.finder.GetAccount(context.TODO(), id)

	s.Nil(account, "Finder.GetAccount should return nil as a first argument, when storage does not hold the account.")
	s.ErrorIsf(err, ErrAccountNotFound, "Finder.GetAccount should return ErrAccountNotFound, when storage does not hold the account, instead of: %s.", err)
}

func (s *FinderTestSuite) TestGetAccountWithAccountInStorage() {
	account, _ := entities.NewAccount("797834131", "Str0ngPassword")
	s.storage.GetByIDFunc = func(ctx context.Context, id uuid.UUID) (*entities.Account, error) {
		if id == account.ID() {
			return account, nil
		}
		return nil, ErrStorageNotFound
	}

	existingAccount, err := s.finder.GetAccount(context.TODO(), account.ID())

	s.Nilf(err, "Finder.GetAccount should return nil as error, if the account exists, instead of: %s", err)
	s.Equal(existingAccount, account, "Finder.GetAccount should return unmodified account from storage")
}

func (s *FinderTestSuite) TestGetAccountWithUnexpectedErroInStorage() {
	s.storage.GetByIDFunc = func(ctx context.Context, id uuid.UUID) (*entities.Account, error) {
		return nil, errors.New("unexpected error from the storage")
	}

	account, err := s.finder.GetAccount(context.TODO(), uuid.New())

	s.NotNil(err, "When an error returned from the storage, Finder.GetAccount should return an error")
	s.ErrorIsf(err, ErrInternalServerError, "When unexpected error returned from the storage, Finder.GetAccount should return ErrInternalServerErr, instead of: %s", err)

	s.Nil(account, "When an error occurred in the storage, Finder.GetAccount should return nil as an account")
}

func (s *FinderTestSuite) TestGetAccountByPhoneWhenAccountDoesNotExistsInTheStorage() {
	s.storage.GetByPhoneFunc = func(ctx context.Context, phone string) (*entities.Account, error) {
		return nil, ErrStorageNotFound
	}

	account, err := s.finder.GetAccountByPhone(context.TODO(), "452314013")

	s.Nil(account, "Finder.GetAccountByPhone should return nil as a first argument, when storage does not hold the account.")
	s.ErrorIsf(err, ErrAccountNotFound, "Finder.GetAccountByPhone should return ErrAccountNotFound, when storage does not hold the account, instead of: %s.", err)
}

func (s *FinderTestSuite) TestGetAccountByPhoneWithAccountInStorage() {
	account, _ := entities.NewAccount("797834131", "Str0ngPassword")
	s.storage.GetByPhoneFunc = func(ctx context.Context, phone string) (*entities.Account, error) {
		if phone == account.Phone() {
			return account, nil
		}
		return nil, ErrStorageNotFound
	}

	existingAccount, err := s.finder.GetAccountByPhone(context.TODO(), account.Phone())

	s.Nilf(err, "Finder.GetAccountByPhone should return nil as error, if the account exists, instead of: %s", err)
	s.Equal(existingAccount, account, "Finder.GetAccountByPhone should return unmodified account from storage")
}

func (s *FinderTestSuite) TestGetAccountByPhoneWithUnexpectedErroInStorage() {
	s.storage.GetByPhoneFunc = func(ctx context.Context, phone string) (*entities.Account, error) {
		return nil, errors.New("unexpected error from the storage")
	}

	account, err := s.finder.GetAccountByPhone(context.TODO(), "3124413123")

	s.NotNil(err, "When an error returned from the storage, Finder.GetAccountByPhone should return an error")
	s.ErrorIsf(err, ErrInternalServerError, "When unexpected error returned from the storage, Finder.GetAccountByPhone should return ErrInternalServerErr, instead of: %s", err)

	s.Nil(account, "When an error occurred in the storage, Finder.GetAccountByPhone should return nil as an account")
}

func TestFinder(t *testing.T) {
	suite.Run(t, new(FinderTestSuite))
}
