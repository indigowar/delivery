package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"github.com/indigowar/delivery/internal/entities"
	"github.com/indigowar/delivery/internal/usecases/auth"
)

type sessionStorage struct {
	client *redis.Client
}

func (s *sessionStorage) GetByID(ctx context.Context, id uuid.UUID) (*entities.Session, error) {
	panic("not implemented")
}

func (s *sessionStorage) GetByToken(ctx context.Context, token entities.SessionToken) (*entities.Session, error) {
	panic("not implemented")
}

func (s *sessionStorage) Add(ctx context.Context, session *entities.Session) (*entities.Session, error) {
	panic("not implemented")
}

func (s *sessionStorage) Remove(ctx context.Context, token string) error {
	panic("not implemented")
}

func (s *sessionStorage) Update(ctx context.Context, session *entities.Session) error {
	panic("not implemented")
}

// insertSession inserts session into redis
func (s *sessionStorage) insertSession(ctx context.Context, sessionId string, session *entities.Session) error {
	// insert a record into redis
	if err := s.client.HSet(ctx, sessionId, "account", session.AccountID, "token", session.Token).Err(); err != nil {
		return err
	}

	// set when the session expires
	if err := s.client.Expire(ctx, sessionId, session.Duration).Err(); err != nil {
		return err
	}

	// add indices for search

	if err := s.client.Set(ctx, fmt.Sprintf("token:%s", session.Token), sessionId, session.Duration).Err(); err != nil {
		return err
	}

	if err := s.client.Set(ctx, fmt.Sprintf("account:%s", session.AccountID.String()), sessionId, session.Duration).Err(); err != nil {
		return err
	}

	return nil
}

func (s *sessionStorage) getSession(ctx context.Context, sessionId string) (*entities.Session, error) {
	sessionData, err := s.client.HGetAll(ctx, sessionId).Result()
	if err != nil {
		return nil, err
	}

	duration, err := s.client.ExpireTime(ctx, sessionId).Result()
	if err != nil {
		return nil, err
	}

	id, _ := uuid.Parse(sessionData["account"])

	return &entities.Session{
		AccountID: id,
		Token:     entities.SessionToken(sessionData["token"]),
		Duration:  duration,
	}, nil
}

func (s *sessionStorage) generateSessionID(ctx context.Context, account uuid.UUID) string {
	return fmt.Sprintf("session_%s", account.String())
}

func NewSessionStorage(client *redis.Client, expTime time.Duration) auth.SessionStorage {
	return &sessionStorage{
		client: client,
	}
}
