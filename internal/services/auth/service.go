package auth

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/indigowar/delivery/internal/domain"
)

var (
	ErrCredentialsAreInvalid = errors.New("credentials are invalid")
	ErrAlreadyInUse          = errors.New("those credentials are already in use")
)

type Service interface {
	Login(ctx context.Context, phone string, password string) (uuid.UUID, error)
	Register(ctx context.Context, phone string, password string) (uuid.UUID, error)
}

type service struct {
	getByPhone domain.GetAccountByPhoneUseCase
	create     domain.CreateAccountUseCase
}

func NewService(getByPhone domain.GetAccountByPhoneUseCase, create domain.CreateAccountUseCase) Service {
	return &service{
		getByPhone: getByPhone,
		create:     create,
	}
}

func (svc *service) Login(ctx context.Context, phone string, password string) (uuid.UUID, error) {
	account, err := svc.getByPhone.GetByPhone(ctx, phone)
	if err != nil {
		if errors.Is(err, domain.ErrAccountNotFound) {
			// this account does not exists,
			// credentials are invalid
			return uuid.UUID{}, ErrCredentialsAreInvalid
		}
		// if some other error was thrown, it means that internal error happened.
		return uuid.UUID{}, errors.New("internal error")
	}

	if !account.HasEqualPassword(password) {
		return uuid.UUID{}, ErrCredentialsAreInvalid
	}

	return account.ID(), nil
}

func (svc *service) Register(ctx context.Context, phone string, password string) (uuid.UUID, error) {
	account, err := domain.NewAccount(phone, password)
	if err != nil {
		log.Println(err)
		return uuid.UUID{}, ErrCredentialsAreInvalid
	}

	if err = svc.create.Create(ctx, &account); err != nil {
		if errors.Is(err, domain.ErrAccountAlreadyExists) {
			return uuid.UUID{}, ErrAlreadyInUse
		}
		// if some other error was thrown, it means that internal error happened.
		return uuid.UUID{}, errors.New("internal error")
	}
	return account.ID(), nil
}
