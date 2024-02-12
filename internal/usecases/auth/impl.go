package auth

import (
	"context"
	"errors"
	"time"

	"github.com/indigowar/delivery/internal/entities"
)

type ServiceImplementation struct {
	tokenGenerator ShortLiveTokenGenerator
	storage        SessionStorage
	validator      CredentialsValidator

	secret              []byte
	sessionLifetime     time.Duration
	accessTokenLifetime time.Duration
}

func (svc *ServiceImplementation) StartSession(ctx context.Context, phone string, password string) (*TokenResult, error) {
	id, err := svc.validator.Validate(ctx, phone, password)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	_, err = svc.storage.GetByID(ctx, id)
	if err == nil {
		return nil, ErrSessionAlreadyExists
	}

	if !errors.Is(err, ErrSessionNotFoundInStorage) {
		return nil, ErrInternalServerError
	}

	session, err := svc.addSession(ctx, entities.NewSession(id, svc.sessionLifetime))
	if err != nil {
		return nil, err
	}

	return svc.makeTokenResult(session)
}

func (svc *ServiceImplementation) ExtendSession(ctx context.Context, token entities.SessionToken) (*TokenResult, error) {
	session, err := svc.storage.GetByToken(ctx, token)
	if err != nil {
		if errors.Is(err, ErrSessionNotFoundInStorage) {
			return nil, ErrSessionDoesNotExists
		}

		return nil, ErrInternalServerError
	}

	session = entities.NewSession(session.AccountID, svc.sessionLifetime)

	if err := svc.storage.Update(ctx, session); err != nil {
		if errors.Is(err, ErrSessionDoesNotExists) {
			return nil, ErrSessionDoesNotExists
		}

		return nil, ErrInternalServerError
	}

	return svc.makeTokenResult(session)
}

func (svc *ServiceImplementation) EndSession(ctx context.Context, sessionToken entities.SessionToken) error {
	if err := svc.storage.Remove(ctx, string(sessionToken)); err != nil {
		if errors.Is(err, ErrSessionNotFoundInStorage) {
			return ErrSessionDoesNotExists
		}

		return ErrInternalServerError
	}

	return nil
}

func (svc *ServiceImplementation) GetAccessToken(ctx context.Context, sessionToken entities.SessionToken) (string, error) {
	session, err := svc.storage.GetByToken(ctx, sessionToken)
	if err != nil {
		if errors.Is(err, ErrSessionNotFoundInStorage) {
			return "", ErrSessionDoesNotExists
		}

		return "", ErrInternalServerError
	}

	return svc.createAccessToken(session)
}

func (svc *ServiceImplementation) makeTokenResult(session *entities.Session) (*TokenResult, error) {
	token, _ := svc.createAccessToken(session)

	return &TokenResult{
		Session:        session,
		ShortLiveToken: token,
	}, nil
}

func (svc *ServiceImplementation) createAccessToken(session *entities.Session) (string, error) {
	return svc.tokenGenerator.Generate(TokenPayload{
		AccountID: session.AccountID,
		Issuer:    "auth_svc",
		Duration:  svc.accessTokenLifetime,
		Key:       svc.secret,
	})
}

func (svc *ServiceImplementation) addSession(ctx context.Context, session *entities.Session) (*entities.Session, error) {
	session, err := svc.storage.Add(ctx, session)
	if err != nil {
		if errors.Is(err, ErrSessionAlreadyInStorage) {
			return nil, ErrSessionAlreadyExists
		}

		return nil, ErrInternalServerError
	}

	return session, nil
}

type CreationOption func(i *ServiceImplementation)

func WithStorage(storage SessionStorage) CreationOption {
	return func(svc *ServiceImplementation) {
		svc.storage = storage
	}
}

func WithTokenGenerator(generator ShortLiveTokenGenerator) CreationOption {
	return func(svc *ServiceImplementation) {
		svc.tokenGenerator = generator
	}
}

func WithCredentialsValidator(validator CredentialsValidator) CreationOption {
	return func(svc *ServiceImplementation) {
		svc.validator = validator
	}
}

func WithSecret(secret []byte) CreationOption {
	return func(svc *ServiceImplementation) {
		svc.secret = secret
	}
}

func WithSessionLifetime(lifetime time.Duration) CreationOption {
	return func(svc *ServiceImplementation) {
		svc.sessionLifetime = lifetime
	}
}

func WithAccessTokenLifetime(lifetime time.Duration) CreationOption {
	return func(svc *ServiceImplementation) {
		svc.accessTokenLifetime = lifetime
	}
}

func NewServiceImplementation(options ...CreationOption) *ServiceImplementation {
	impl := &ServiceImplementation{}

	for _, option := range options {
		option(impl)
	}

	return impl
}
