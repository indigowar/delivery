package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Validator struct {
	secret []byte
}

func NewValidator(secret []byte) *Validator {
	return &Validator{
		secret: secret,
	}
}

func (v *Validator) Validate(tokenString string) (uuid.UUID, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return v.secret, nil
	})

	if err != nil {
		return uuid.UUID{}, err
	}

	if !token.Valid {
		return uuid.UUID{}, errors.New("token is not valid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.UUID{}, errors.New("failed to extract claims")
	}

	uuidStr, ok := claims["sub"].(string)
	if !ok {
		return uuid.UUID{}, errors.New("subject is missing or invalid")
	}

	parsedUUID, err := uuid.Parse(uuidStr)
	if err != nil {
		return uuid.UUID{}, errors.New("failed to parse UUID")
	}

	return parsedUUID, nil
}
