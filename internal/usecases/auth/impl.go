package auth

import (
	"context"
	"time"

	"github.com/indigowar/delivery/internal/entities"
)

type impl struct {
	storage        SessionStorage
	tokenGenerator ShortLiveTokenGenerator

	secret   []byte
	duration time.Duration
}

func (svc *impl) StartSession(ctx context.Context, phone string, password string) (*TokenResult, error) {
	// todo: implement
	panic("unimplemented")
}

func (svc *impl) ExtendSession(ctx context.Context, sessionToken entities.SessionToken) (*TokenResult, error) {
	// todo: implement
	panic("unimplemented")
}

func (svc *impl) EndSession(ctx context.Context, sessionToken entities.SessionToken) (*TokenResult, error) {
	// todo: implement
	panic("unimplemented")
}

func (svc *impl) GetAccessToken(ctx context.Context, sessionToken entities.SessionToken) (string, error) {
	// todo: implement
	panic("unimplemented")
}

func NewService() Service {
	return &impl{}
}
