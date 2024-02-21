package entities

import (
	"net/url"

	"github.com/google/uuid"
)

type Category struct {
	ID     uuid.UUID
	Name   string
	Image  *url.URL
	Dishes []uuid.UUID
}
