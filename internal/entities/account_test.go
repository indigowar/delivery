package entities

import (
	"net/mail"
	"testing"

	"github.com/stretchr/testify/suite"
)

type AccountTestSuite struct {
	suite.Suite

	account *Account
}

func (s *AccountTestSuite) SetupTest() {
	s.account = &Account{}
}

func (s *AccountTestSuite) TestSetPhone() {
	validPhoneNumbers := []string{
		"1234567890",
		"123453311213",
		"213123213425761",
		"12345",
		"3560",
	}

	for _, phone := range validPhoneNumbers {
		err := s.account.SetPhone(phone)
		s.Nilf(err, "SetPhone returns an error(%s), when given valid phone number: %s!", err, phone)
	}

	invalidPhoneNumbers := []string{
		"a text ",
		"-123312-4323",
		"11111111111111111",
		"12341cv12",
		"fv0213_da",
		"2131@",
		"#123543",
		" 21312312312 ",
	}

	for _, phone := range invalidPhoneNumbers {
		err := s.account.SetPhone(phone)
		s.NotNilf(err, "SetPhone should not return  nil, when give invalid phone number: %s!", phone)

		s.ErrorIsf(err, ErrProvidedDataIsInvalid, "SetPhone should return ErrProvidedDataIsInvalid, when the invalid phone number given, instead of: %w!", err)
	}
}

func (s *AccountTestSuite) TestSetPassword() {
	validPasswords := []string{
		"c0mplicatedP@ssword",
		"simplePass123",
		"#123fsg321Daedr11dfcdxza@",
		"enoughtComplicated44",
		"DbMA123-d)231",
	}

	for _, word := range validPasswords {
		err := s.account.SetPassword(word)
		s.Nilf(err, "SetPassword returns an error(%s), when given valid password: %s!", err, word)
	}

	invalidPasswords := []string{
		"",
		"  ",
		" ",
		"password",
		"simple",
		"qwerty",
		"123456",
		"justatext",
		"small",
		"text with spaces",
		"NoDigitsOrSpecial",
	}

	for _, word := range invalidPasswords {
		err := s.account.SetPassword(word)

		s.NotNilf(err, "SetPassword should not return nil, when give invalid password: %s!", word)

		s.ErrorIsf(err, ErrProvidedDataIsInvalid, "SetPassword should return ErrProvidedDataIsInvalid, when the invalid password given, instead of: %w!", err)
	}
}

func (s *AccountTestSuite) TestSetEmail() {
	var invalidEmail *mail.Address = nil
	validEmail, _ := mail.ParseAddress("example@mail.com")

	err := s.account.SetEmail(validEmail)
	s.Nilf(err, "SetEmail return nil, when valid email is provided, instead of: %s!", err)

	err = s.account.SetEmail(invalidEmail)
	s.NotNil(err, "SetEmail should not return nil, when invalid email is provided!")

	s.ErrorIsf(err, ErrProvidedDataIsInvalid, "SetEmail should return ErrProvidedDataIsInvalid, when invalid email provided, instead of: %w!", err)
}

func (s *AccountTestSuite) TestSetFirstName() {
	validNames := []string{
		"Max",
		"Alexander",
		"Vladislav",
		"Mike",
		"John",
		"Jessy",
		"Nick",
	}

	for _, name := range validNames {
		err := s.account.SetFirstName(name)
		s.Nilf(err, "SetFirstName should return nil, when valid name is provided, instead of %s!", err)
	}

	invalidNames := []string{
		"M1231",
		"Some one other",
		"X",
		"",
		"@3wadsc@3WDAS",
	}

	for _, name := range invalidNames {
		err := s.account.SetFirstName(name)
		s.NotNilf(err, "SetFirstName should return an error, invalid value is provided: %s!", name)

		s.ErrorIsf(err, ErrProvidedDataIsInvalid, "SetFirstName should return ErrProvidedDataIsInvalid instead of: %s!", err)
	}
}

func (s *AccountTestSuite) TestSetSurname() {
	validSurnames := []string{
		"Doe",
		"Smith",
		"Johnson",
		"Williams",
		"Brown",
		"Jones",
		"Miller",
	}

	for _, surname := range validSurnames {
		err := s.account.SetSurname(surname)
		s.Nilf(err, "SetSurname should return nil when a valid surname is provided, instead of %s!", err)
	}

	invalidSurnames := []string{
		"1234",
		"Some one other",
		"X",
		"",
		"@3wadsc@3WDAS",
	}

	for _, surname := range invalidSurnames {
		err := s.account.SetSurname(surname)
		s.NotNilf(err, "SetSurname should return an error for an invalid value: %s!", surname)
		s.ErrorIsf(err, ErrProvidedDataIsInvalid, "SetSurname should return ErrProvidedDataIsInvalid for an invalid value: %s!", surname)
	}
}

func (s *AccountTestSuite) TestName() {
	firstname := "Alex"
	surname := "Doe"

	s.account.firstName = &firstname
	s.account.surname = nil

	s.Equal(s.account.Name(), *s.account.firstName, "The Account.Name() should return the first name, if it is defined and surname is not.")
	s.True(s.account.HasName(), "Account.HasName() should return true, if firstname is defined")

	s.account.surname = &surname
	expected := *s.account.firstName + " " + *s.account.surname
	s.Equalf(s.account.Name(), expected, "The Account.Name() should return `FirstName Surname`, if they are both defined")
	s.True(s.account.HasName(), "Account.HasName() should return true, if firstname and surname are defined")

	s.account.surname = nil
	s.account.firstName = nil
	s.False(s.account.HasName(), "Account.HasName() should return false, if firstname and surname are undefined")
}

func TestAccountTestSuite(t *testing.T) {
	suite.Run(t, new(AccountTestSuite))
}
