package domain

import (
	"errors"
	"fmt"
	"net/mail"
	"net/url"
	"regexp"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	id       uuid.UUID
	phone    string
	password string

	// not required fields, can be empty

	email      *mail.Address
	firstName  *string
	surname    *string
	pictureURL *url.URL
}

func NewAccount(phone string, password string) (Account, error) {
	a := Account{
		id: uuid.New(),
	}

	if err := a.SetPhone(phone); err != nil {
		return Account{}, err
	}

	if err := a.SetPassword(password); err != nil {
		return Account{}, err
	}

	return a, nil
}

func (a *Account) ID() uuid.UUID {
	return a.id
}

func (a *Account) Phone() string {
	return a.phone
}

func (a *Account) SetPhone(value string) error {
	a.phone = value
	// todo: add validation
	return nil
}

func (a *Account) ValidatePassword(value string) bool {
	return comparePasswordToReal(a.password, value)
}

func (a *Account) SetPassword(password string) error {
	if err := validatePasswordValue(password); err != nil {
		return err
	}

	hashed, err := hashPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	a.password = hashed

	return nil
}

func (a *Account) HasEmail() bool {
	return a.email != nil
}

func (a *Account) Email() *mail.Address {
	return a.email
}

func (a *Account) SetEmail(addr *mail.Address) {
	a.email = addr
}

func (a *Account) HasName() bool {
	return a.firstName != nil || a.surname != nil
}

func (a *Account) SetFirstName(value string) error {
	if len(value) < 3 {
		return errors.New("given value for first name is too small")
	}
	a.firstName = &value
	return nil
}

func (a *Account) SetSurname(value string) error {
	if len(value) < 3 {
		return errors.New("given value for surname is too small")
	}
	a.surname = &value
	return nil
}

func (a *Account) GetName() string {
	name := ""

	if a.firstName != nil {
		name += *a.firstName
	}

	if a.surname != nil && a.firstName != nil {
		name += " " + *a.surname
	} else if a.surname != nil {
		name += *a.surname
	}

	return name
}

func (a *Account) HasPicture() bool {
	return a.pictureURL != nil
}

func (a *Account) SetPictureURL(value *url.URL) {
	a.pictureURL = value
}

func (a *Account) GetPictureURL() *url.URL {
	return a.pictureURL
}

// validatePasswordValue - returns an error, that indicates
// what criteria given password value doesn't fit, or nil.
func validatePasswordValue(value string) error {
	if len(value) < 6 {
		return errors.New("password value is too small")
	}

	if ok, _ := regexp.MatchString(`[A-Z]`, value); !ok {
		return errors.New("password must contain at least one uppercase letter")
	}

	if ok, _ := regexp.MatchString(`[a-z]`, value); !ok {
		return errors.New("password must contain at least one lowercase letter")
	}

	if ok, _ := regexp.MatchString(`[0-9]`, value); !ok {
		return errors.New("password must contain at least one digit")
	}

	return nil
}

func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func comparePasswordToReal(real string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(real), []byte(password)) == nil
}
