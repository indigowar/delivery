package entities

import (
	"errors"
	"net/url"

	"github.com/google/uuid"
)

type Category struct {
	id         uuid.UUID
	name       string
	restaurant uuid.UUID
	image      *url.URL
	dishes     map[uuid.UUID]struct{}
}

func NewCategory(name string, restaurant uuid.UUID, image *url.URL, dishes ...uuid.UUID) (*Category, error) {
	return nil, nil
}

func (c Category) ID() uuid.UUID {
	return c.id
}

func (c Category) Name() string {
	return c.name
}

func (c *Category) SetName(value string) error {
	if len(value) < 3 {
		return errors.New("category name should contain at least 3 letters")
	}

	c.name = value

	return nil
}

func (c Category) Restaurant() uuid.UUID {
	return c.restaurant
}

func (c Category) Image() *url.URL {
	return c.image
}

func (c Category) HasDish(id uuid.UUID) bool {
	for v := range c.dishes {
		if v == id {
			return true
		}
	}
	return false
}

func (c *Category) AddDish(id uuid.UUID) {
	c.dishes[id] = struct{}{}
}

func (c *Category) DeleteDish(id uuid.UUID) {
	delete(c.dishes, id)
}

func (c *Category) GetDishes() []uuid.UUID {
	ids := make([]uuid.UUID, len(c.dishes))
	i := 0
	for v := range c.dishes {
		ids[i] = v
		i++
	}

	return ids
}
