package redis

import (
	"context"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"github.com/indigowar/delivery/internal/entities"
	connector "github.com/indigowar/delivery/pkg/redis"
)

type Storage struct {
	client *redis.Client
}

// GetByID implements auth.SessionStorage
func (s *Storage) GetByID(ctx context.Context, id uuid.UUID) (*entities.Session, error) {
	session, err := getByAccountID(ctx, s.client, id)
	if err != nil {
		return nil, err
	}

	return toEntity(session), nil
}

// GetByToken implements auth.SessionStorage
func (s *Storage) GetByToken(ctx context.Context, token entities.SessionToken) (*entities.Session, error) {
	session, err := getByToken(ctx, s.client, string(token))
	if err != nil {
		return nil, err
	}

	return toEntity(session), nil
}

// Add implements auth.SessionStorage
func (s *Storage) Add(ctx context.Context, session *entities.Session) (*entities.Session, error) {
	data := fromEntity(session)
	data.ID = uuid.New()

	if err := data.Save(ctx, s.client); err != nil {
		return nil, err
	}

	return toEntity(data), nil
}

// Remove implements auth.SessionStorage
func (s *Storage) Remove(ctx context.Context, token string) error {
	session, err := getByToken(ctx, s.client, token)
	if err != nil {
		return err
	}

	return session.Delete(ctx, s.client)
}

// Update implements auth.SessionStorage
func (s *Storage) Update(ctx context.Context, session *entities.Session) error {
	existingAccount, err := getByAccountID(ctx, s.client, session.AccountID)
	if err != nil {
		return err
	}

	return existingAccount.Update(ctx, s.client, session)
}

func (s *Storage) Close() error {
	return s.client.Close()
}

func NewStorage(host string, port string, password string) (*Storage, error) {
	client, err := connector.ConnectToRedis(host, port, password)
	if err != nil {
		return nil, err
	}

	return &Storage{
		client: client,
	}, nil
}
