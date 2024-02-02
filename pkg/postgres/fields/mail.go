package fields

import (
	"database/sql/driver"
	"fmt"
	"net/mail"
)

type MailAdress mail.Address

func (m *MailAdress) Scan(value interface{}) error {
	if value == nil {
		*m = MailAdress{}
		return nil
	}

	s, ok := value.(string)
	if !ok {
		return fmt.Errorf("unexpected type %T, expected string", value)
	}

	address, err := mail.ParseAddress(s)
	if err != nil {
		return err
	}

	*m = MailAdress(*address)

	return nil
}

func (m MailAdress) Value() (driver.Value, error) {
	email := mail.Address(m)
	return email.String(), nil
}
