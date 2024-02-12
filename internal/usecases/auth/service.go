package auth

import (
	"context"
	"errors"

	"github.com/indigowar/delivery/internal/entities"
)

// List of errors that can be returned from auth.Service
var (
	// ErrInvalidCredentials is returned, when the given phone and password to
	// Service.StartSession are invalid:
	// - values is invalid.
	// - there is no account with given phone.
	// - the password isn't right for given account.
	ErrInvalidCredentials = errors.New("provided credentials are invalid")
	// ErrSessionAlreadyExists is returned, when given session token
	// to ExtendSession or EndSession does not belong to a session.
	ErrSessionDoesNotExists = errors.New("session does not exists")
	// ErrSessionAlreadyExists is returned, when StartSession receives
	// a phone number for a user, that is already has a session.
	ErrSessionAlreadyExists = errors.New("session already exists")
	// ErrInternalServerError is returned, when an error occurred in the system.
	ErrInternalServerError = errors.New("internal server error")
)

// TokenResult is a simple structure, that holds a short live token and the session itself.
type TokenResult struct {
	ShortLiveToken string
	Session        *entities.Session
}

// auth.Service is the interface of the auth microservice
type Service interface {
	// StartSession - starts a new session for a user with given phone,
	// if the password is valid and there is no active session already.
	//
	// Returns a short live token, that is used for access, session token(long live)
	// and id of a user.
	StartSession(ctx context.Context, phone string, password string) (*TokenResult, error)

	// ExtendSession - extends current session and returns a new tokens of it.
	ExtendSession(ctx context.Context, sessionToken entities.SessionToken) (*TokenResult, error)

	// EndSession - ends currently active user session
	EndSession(ctx context.Context, sessionToken entities.SessionToken) error

	// GetAccessToken - creates and returns a new short-live access token for user.
	GetAccessToken(ctx context.Context, sessionToken entities.SessionToken) (string, error)
}
