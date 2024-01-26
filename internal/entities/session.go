package entities

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/google/uuid"
)

type SessionToken string

type Session struct {
	AccountID uuid.UUID
	Token     SessionToken
}

func NewSession(accountID uuid.UUID) *Session {
	token, _ := generateUniqueString(32)

	return &Session{
		AccountID: accountID,
		Token:     SessionToken(token),
	}
}

func generateUniqueString(length int) (string, error) {
	byteLength := (length * 6) / 8
	if (length*6)%8 != 0 {
		byteLength++
	}

	bytes := make([]byte, byteLength)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	uniqueString := base64.URLEncoding.EncodeToString(bytes)

	if len(uniqueString) > length {
		uniqueString = uniqueString[:length]
	}

	return uniqueString, nil
}
