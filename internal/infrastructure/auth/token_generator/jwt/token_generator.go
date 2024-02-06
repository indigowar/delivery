package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/indigowar/delivery/internal/usecases/auth"
)

type TokenGenerator struct{}

func NewTokenGenerator() *TokenGenerator {
	return &TokenGenerator{}
}

func (tg *TokenGenerator) Generate(payload auth.TokenPayload) (string, error) {
	expTime := time.Now().Add(payload.Duration)

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.RegisteredClaims{
		Issuer:    payload.Issuer,
		Subject:   payload.AccountID.String(),
		ExpiresAt: jwt.NewNumericDate(expTime),
	})

	return token.SignedString(payload.Key)
}
