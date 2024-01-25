package entities

import (
	"errors"
	"fmt"
	"net/mail"
	"net/url"
	"regexp"

	"github.com/google/uuid"
)

var (
	ErrProvidedDataIsInvalid = errors.New("provided data is invalid")
)

type Account struct {
	id uuid.UUID

	phone    string
	password string

	email           *mail.Address
	firstName       *string
	surname         *string
	profileImageUrl *url.URL
}

func NewAccount(phone string, password string) (*Account, error) {
	a := &Account{
		id: uuid.New(),
	}

	if err := a.SetPhone(phone); err != nil {
		return nil, err
	}

	if err := a.SetPassword(password); err != nil {
		return nil, err
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
	if err := a.ValidatePhoneNumberValue(value); err != nil {
		return err
	}

	a.phone = value

	return nil
}

func (a *Account) SetPassword(value string) error {
	if !a.ValidatePasswordValue(value) {
		return ErrProvidedDataIsInvalid
	}

	a.password = value

	return nil
}

func (a *Account) HasEmail() bool {
	return a.email != nil
}

func (a *Account) Email() *mail.Address {
	return a.email
}

func (a *Account) SetEmail(addr *mail.Address) error {
	if addr == nil {
		return fmt.Errorf("%w: email address should not be nil", ErrProvidedDataIsInvalid)
	}

	a.email = addr

	return nil
}

func (a *Account) HasName() bool {
	return a.firstName != nil || a.surname != nil
}

func (a *Account) Name() string {
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

func (a *Account) SetFirstName(value string) error {
	if len(value) < 3 {
		return fmt.Errorf("%w: first name should have at least 3 characters", ErrProvidedDataIsInvalid)
	}

	if !containsOnlyLetters(value) {
		return fmt.Errorf("%w: first name should contain only letters", ErrProvidedDataIsInvalid)
	}

	a.firstName = &value

	return nil
}

func (a *Account) SetSurname(value string) error {
	if len(value) < 3 {
		return fmt.Errorf("%w: surname should have at least 3 characters", ErrProvidedDataIsInvalid)
	}

	if !containsOnlyLetters(value) {
		return fmt.Errorf("%w: surname should contain only letters", ErrProvidedDataIsInvalid)
	}

	a.surname = &value

	return nil
}

func (a *Account) HasProfileImageUrl() bool {
	return a.profileImageUrl != nil
}

func (a *Account) ProfileImageUrl() *url.URL {
	return a.profileImageUrl
}

func (a *Account) SetProfileImageUrl(link *url.URL) error {
	if link == nil {
		return fmt.Errorf("%w: the profile image url should not be nil", ErrProvidedDataIsInvalid)
	}

	a.profileImageUrl = link

	return nil
}

func (Account) ValidatePhoneNumberValue(value string) error {
	pattern := regexp.MustCompile(`^\d{1,15}$`)

	if !pattern.MatchString(value) {
		return ErrProvidedDataIsInvalid
	}
	return nil
}

func (Account) ValidatePasswordValue(password string) bool {
	conditions := 0

	hasLower := regexp.MustCompile(`[a-z]`).MatchString
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString
	hasDigit := regexp.MustCompile(`\d`).MatchString
	hasSpecial := regexp.MustCompile(`[^\w\s]`).MatchString

	if hasLower(password) {
		conditions++
	}
	if hasUpper(password) {
		conditions++
	}
	if hasDigit(password) {
		conditions++
	}
	if hasSpecial(password) {
		conditions++
	}

	return conditions >= 3 && len(password) >= 8
}

func containsOnlyLetters(value string) bool {
	pattern := regexp.MustCompile(`^[a-zA-Z]+$`)
	return pattern.MatchString(value)
}
