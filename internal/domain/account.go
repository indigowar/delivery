package domain

import (
	"errors"
	"log"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	id       uuid.UUID
	phone    string
	password string
}

func (a *Account) ID() uuid.UUID {
	return a.id
}

func (a *Account) Phone() string {
	return a.phone
}

func (a *Account) SetPhone(value string) error {
	if !ValidatePhoneNumber(value) {
		return errors.New("phone value is invalid")
	}
	a.phone = value
	return nil
}

func (a *Account) Password() string {
	return a.password
}

func (a *Account) HasEqualPassword(value string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(value), []byte(a.password))
	return err == nil
}

func (a *Account) SetPassword(value string) error {
	if !ValidatePassword(value) {
		return errors.New("password value is invalid")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(value), 14)
	if err != nil {
		log.Printf("failed to generate hash from password: %s\n", err)
		return errors.New("crypto error")
	}

	a.password = string(bytes)
	return nil
}

func NewAccount(phone string, password string) (Account, error) {
	account := Account{
		id: uuid.New(),
	}

	if err := account.SetPhone(phone); err != nil {
		return Account{}, err
	}

	if err := account.SetPassword(password); err != nil {
		return Account{}, err
	}

	return account, nil
}

func ValidatePhoneNumber(phone string) bool {
	phoneRegex := regexp.MustCompile(`\d{10,}`)
	return phoneRegex.MatchString(phone)
}

func ValidatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return false
	}

	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		return false
	}

	if !strings.ContainsAny(password, "0123456789") {
		return false
	}

	specialCharRegex := regexp.MustCompile(`[!@#$%^&*]`)
	if !specialCharRegex.MatchString(password) {
		return false
	}

	return true
}
