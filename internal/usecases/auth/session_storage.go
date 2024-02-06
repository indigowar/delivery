package auth

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/indigowar/delivery/internal/entities"
)

//go:generate moq -out session_storage_moq_test.go . SessionStorage

var (
	StorageErrSessionNotFound      = errors.New("session is not found")
	StorageErrSessionAlreadyExists = errors.New("session is already exists")
)

type SessionStorage interface {
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Session, error)
	GetByToken(ctx context.Context, token entities.SessionToken) (*entities.Session, error)
	Add(ctx context.Context, session *entities.Session) (*entities.Session, error)
	Remove(ctx context.Context, token string) error
	Update(ctx context.Context, session *entities.Session) error
}
