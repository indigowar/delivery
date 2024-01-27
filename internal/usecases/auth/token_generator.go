package auth

import (
	"time"

	"github.com/google/uuid"
)

//go:generate moq -out token_generator_moq_test.go . ShortLiveTokenGenerator

type TokenPayload struct {
	AccountID    uuid.UUID
	AccountPhone string
	Issuer       string
	Duration     time.Duration
	Key          []byte
}

type ShortLiveTokenGenerator interface {
	Generate(payload TokenPayload) (string, error)
}
