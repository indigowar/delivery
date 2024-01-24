package domain

import (
	"context"
	"errors"
	"mime/multipart"
	"net/mail"

	"github.com/google/uuid"
)

var (
	// ErrAccountNotFound is been returned when the account that is been searched is not found.
	ErrAccountNotFound = errors.New("account is not found")
	// ErrAccountAlreadyExists is been returned when the phone number is used,
	// that belong to existing another account.
	ErrAccountAlreadyExists = errors.New("account is already exists with given credentials")
	// ErrInvalidCredentials is been returned when the credentials sended to service are invalid.
	ErrInvalidCredentials = errors.New("given credentials are invalid")
	// ErrEmailIsAlreadyInUse is been returned when the e-mail is already used
	// by another account.
	ErrEmailIsAlreadyInUse = errors.New("given email is already in use by some other account")
	// ErrInternalServerError is been returned when something unexpected happened in the system
	ErrInternalServerError = errors.New("internal server error")
)

type Service interface {
	// GetByID - returns an account with id from the arguments or an error
	// otherwise returns ErrAccountNotFound or ErrInternalServerError
	GetByID(ctx context.Context, id uuid.UUID) (*Account, error)

	// GetByPhone - returns an account with phone from the arguments
	// otherwise returns ErrAccountNotFound or ErrInternalServerError
	GetByPhone(ctx context.Context, phone string) (*Account, error)

	// ValidateCredentials - returns an id of account with given phone and password
	// otherwise returns ErrInvalidCredentials or ErrInternalServerError
	ValidateCredentials(ctx context.Context, phone string, password string) (uuid.UUID, error)

	// Register - registers a new user in the service with given credentials.
	// Can return ErrInvalidCredentials, ErrAccountAlreadyExists, ErrInternalServerError
	Register(ctx context.Context, phone string, password string) (uuid.UUID, error)

	// AddProfileImage - adds profile image to user account.
	// Can return ErrInternalServerError, ErrAccountNotFound
	AddProfileImage(ctx context.Context, ownerId uuid.UUID, profileImage multipart.File) error

	// AddEmail - sets this email to the account with given id.
	// Can return ErrAccountNotFound, ErrEmailIsAlreadyInUse or ErrInternalServerError
	AddEmail(ctx context.Context, id uuid.UUID, mail *mail.Address) error
}
