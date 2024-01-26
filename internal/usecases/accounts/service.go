package accounts

import (
	"context"
	"errors"
	"net/mail"
	"os"

	"github.com/google/uuid"
	"github.com/indigowar/delivery/internal/entities"
)

// List of errors that can be returned by accounts service,
// any other error is considered to be ErrInternalServerError
var (
	// returned, when the account is not found either for search, or by id
	ErrAccountNotFound = errors.New("account is not found")
	// returned, when register an account on the phone, that already in use
	ErrAccountIsAlreadyExists = errors.New("account is already exists")
	// returned, when provided credentials phone and password are invalid
	ErrInvalidCredentials = errors.New("provided credentials are invalid")
	// returned, when provided data for updating account is invalid
	ErrProvidedDataIsInvalid = errors.New("provided data is invalid")
	// returned, when unexpected error occurred in the service
	ErrInternalServerError = errors.New("internal server error occurred")
)

// Finder - is an interface to the account service, that provides retrieve information about account
type Finder interface {
	GetAccount(ctx context.Context, id uuid.UUID) (*entities.Account, error)
	GetAccountByPhone(ctx context.Context, phone string) (*entities.Account, error)
}

// Registrator - is an interface to the account service, that provides registration of a new account.
type Registrator interface {
	RegisterAccount(ctx context.Context, phone string, password string) (uuid.UUID, error)
}

// ProfileUpdater - is an interface to the account service, that is used to update information in existing account
type ProfileUpdater interface {
	LinkEmailToAccount(ctx context.Context, id uuid.UUID, addr *mail.Address) error
	UpdateFirstName(ctx context.Context, id uuid.UUID, firstName string) error
	UpdateSurname(ctx context.Context, id uuid.UUID, surname string) error
	LoadProfileImage(ctx context.Context, id uuid.UUID, image os.File) error
}

// CredentialsValidator - is an interface to the account service, that is used to validate credentials
type CredentialsValidator interface {
	ValidateCredentials(ctx context.Context, phone string, password string) error
}
