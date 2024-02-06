package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"github.com/indigowar/delivery/internal/entities"
	"github.com/indigowar/delivery/internal/usecases/auth"
)

type session struct {
	ID        uuid.UUID `json:"id"`
	AccountID uuid.UUID `json:"account"`
	Token     string    `json:"token"`
	ExpTime   time.Time `json:"exp_time"`
}

func fromEntity(s *entities.Session) session {
	return session{
		ID:        uuid.UUID{},
		AccountID: s.AccountID,
		Token:     string(s.Token),
		ExpTime:   s.ExpirationTime,
	}
}

func toEntity(s session) *entities.Session {
	return &entities.Session{
		AccountID:      s.AccountID,
		Token:          entities.SessionToken(s.Token),
		ExpirationTime: s.ExpTime,
	}
}

func fromBinary(data []byte) (session, error) {
	s := session{}

	if err := json.Unmarshal(data, &s); err != nil {
		return session{}, err
	}

	return s, nil
}

func toBinary(s session) ([]byte, error) {
	return json.Marshal(s)
}

func (s session) Save(ctx context.Context, client *redis.Client) error {
	binary, err := toBinary(s)
	if err != nil {
		return fmt.Errorf("failed to marshalize: %w", err)
	}

	ttl := time.Until(s.ExpTime)

	_, err = client.TxPipelined(ctx, func(p redis.Pipeliner) error {
		if err := p.SetNX(ctx, s.ID.String(), binary, ttl).Err(); err != nil {
			return err
		}

		if err := p.SetNX(ctx, s.AccountID.String(), s.ID.String(), ttl).Err(); err != nil {
			return auth.StorageErrSessionAlreadyExists
		}

		if err := p.SetNX(ctx, s.Token, s.ID.String(), ttl).Err(); err != nil {
			return auth.StorageErrSessionAlreadyExists
		}

		return nil
	})

	return err
}

func (s session) Delete(ctx context.Context, client *redis.Client) error {
	_, err := client.TxPipelined(ctx, func(p redis.Pipeliner) error {
		if err := p.Del(ctx, s.ID.String(), s.AccountID.String(), s.Token); err != nil {
			return auth.StorageErrSessionNotFound
		}

		return nil
	})

	return err
}

func (s *session) Update(ctx context.Context, client *redis.Client, newEntity *entities.Session) error {
	newData := fromEntity(newEntity)

	binary, err := toBinary(newData)
	if err != nil {
		return fmt.Errorf("failed to marshalize: %w", err)
	}

	ttl := time.Until(newData.ExpTime)

	_, err = client.TxPipelined(ctx, func(p redis.Pipeliner) error {
		if err := p.Set(ctx, s.ID.String(), binary, ttl).Err(); err != nil {
			return fmt.Errorf("failed to update")
		}

		if err := p.Expire(ctx, s.AccountID.String(), ttl).Err(); err != nil {
			return auth.StorageErrSessionNotFound
		}

		if err := p.Del(ctx, s.Token).Err(); err != nil {
			return auth.StorageErrSessionNotFound
		}

		if err := p.SetNX(ctx, string(newData.Token), s.ID.String(), ttl).Err(); err != nil {
			return fmt.Errorf("failed to use token: %w", err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	*s = newData

	return nil
}

func getByAccountID(ctx context.Context, client *redis.Client, accountId uuid.UUID) (session, error) {
	indexResult := client.Get(ctx, accountId.String())
	if err := indexResult.Err(); err != nil {
		return session{}, fmt.Errorf("%w: %w", auth.StorageErrSessionNotFound, err)
	}

	sessionId := indexResult.Val()

	return getByID(ctx, client, sessionId)
}

func getByToken(ctx context.Context, client *redis.Client, token string) (session, error) {
	indexResult := client.Get(ctx, token)
	if err := indexResult.Err(); err != nil {
		return session{}, fmt.Errorf("%w:%w", auth.StorageErrSessionNotFound, err)
	}

	sessionId := indexResult.Val()

	return getByID(ctx, client, sessionId)
}

func getByID(ctx context.Context, client *redis.Client, id string) (session, error) {
	sessionResult := client.Get(ctx, id)
	if err := sessionResult.Err(); err != nil {
		return session{}, fmt.Errorf("%w: %w", auth.StorageErrSessionNotFound, err)
	}

	return fromBinary([]byte(sessionResult.Val()))

}
