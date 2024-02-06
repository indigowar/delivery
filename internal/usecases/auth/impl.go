package auth

import (
	"context"
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
	// todo: implement
	panic("unimplemented")
}

func (svc *ServiceImplementation) ExtendSession(ctx context.Context, sessionToken entities.SessionToken) (*TokenResult, error) {
	// todo: implement
	panic("unimplemented")
}

func (svc *ServiceImplementation) EndSession(ctx context.Context, sessionToken entities.SessionToken) (*TokenResult, error) {
	// todo: implement
	panic("unimplemented")
}

func (svc *ServiceImplementation) GetAccessToken(ctx context.Context, sessionToken entities.SessionToken) (string, error) {
	// todo: implement
	panic("unimplemented")
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
