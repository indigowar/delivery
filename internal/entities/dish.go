package entities

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/google/uuid"
)

type Dish struct {
	id    uuid.UUID
	name  string
	price float64
	image *url.URL

	About       string
	Ingredients []string
}

func NewDish(name string, price float64, image *url.URL) (*Dish, error) {
	dish := &Dish{id: uuid.New(), Ingredients: make([]string, 0)}

	if err := dish.SetName(name); err != nil {
		return nil, err
	}

	if err := dish.SetPrice(price); err != nil {
		return nil, err
	}

	if err := dish.SetImage(image); err != nil {
		return nil, err
	}

	return dish, nil
}

func (d Dish) ID() uuid.UUID {
	return d.id
}

func (d Dish) Name() string {
	return d.name
}

func (d *Dish) SetName(value string) error {
	if err := ValidateDishName(value); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	d.name = value

	return nil
}

func (d Dish) Price() float64 {
	return d.price
}

func (d *Dish) SetPrice(value float64) error {
	if value <= 0 {
		return errors.New("price should be a positive number")
	}

	d.price = value

	return nil
}

func (d Dish) Image() *url.URL {
	return d.image
}

func (d *Dish) SetImage(img *url.URL) error {
	if img == nil {
		return errors.New("dish should have an image")
	}

	d.image = img

	return nil
}

func ValidateDishName(value string) error {
	if len(value) < 3 {
		return errors.New("dish name should have at least 3 letters")
	}

	return nil
}
