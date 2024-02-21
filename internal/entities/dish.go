package entities

import (
	"errors"
	"net/url"

	"github.com/google/uuid"
)

type Dish struct {
	ID          uuid.UUID
	Name        string
	About       string
	Image       *url.URL
	Ingredients []string
	Price       float64
}

func NewDish(name string, price float64) (*Dish, error) {
	d := &Dish{
		ID: uuid.New(),
	}

	if err := ValidateDishName(name); err != nil {
		return nil, err
	}

	if price < 0 {
		return nil, errors.New("provided price for the dish is below zero")
	}

	d.Price = price

	return d, nil
}

func ValidateDishName(value string) error {
	if len(value) < 3 {
		return errors.New("provided name for the dish is invalid")
	}

	return nil
}

func ValidateDishAbout(value string) error {
	panic("not implemented")
}
