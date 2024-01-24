package service

import (
	"context"
	"mime/multipart"
	"net/mail"

	"github.com/google/uuid"

	"github.com/indigowar/delivery/internal/accounts/domain"
)

type svc struct{}

func NewService() domain.Service {
	return &svc{}
}

func (s *svc) GetByID(ctx context.Context, id uuid.UUID) (*domain.Account, error) {
	panic("not implemented")
}

func (s *svc) GetByPhone(ctx context.Context, phone string) (*domain.Account, error) {
	panic("not implemented")
}

func (s *svc) ValidateCredentials(ctx context.Context, phone string, password string) (uuid.UUID, error) {
	panic("not implemented")
}

func (s *svc) Register(ctx context.Context, phone string, password string) (uuid.UUID, error) {
	panic("not implemented")
}

func (s *svc) AddProfileImage(ctx context.Context, ownerId uuid.UUID, profileImage multipart.File) error {
	panic("not implemented")
}

func (s *svc) AddEmail(ctx context.Context, id uuid.UUID, mail *mail.Address) error {
	panic("not implemented")
}
