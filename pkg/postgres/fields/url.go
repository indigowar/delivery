package fields

import (
	"database/sql/driver"
	"fmt"
	"net/url"
)

type URL url.URL

func (u *URL) Scan(value interface{}) error {
	if value == nil {
		*u = URL{}
		return nil
	}

	s, ok := value.(string)
	if !ok {
		return fmt.Errorf("unexpected type %T, expected string", value)
	}

	url, err := url.Parse(s)
	if err != nil {
		return err
	}

	*u = URL(*url)

	return nil
}

func (u URL) Value() (driver.Value, error) {
	url := url.URL(u)
	return url.String(), nil
}
