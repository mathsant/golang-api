package entity

import (
	"errors"
	"time"

	"github.com/mathsant/golang-api/pkg/entity"
)

var (
	ErrIDIsRequired    = errors.New("id is required")
	ErrIDIsInvalid     = errors.New("invalid id")
	ErrNameIsRequired  = errors.New("name is requrired")
	ErrPriceIsRequired = errors.New("price is requrired")
	ErrPriceIsInvalid  = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		ID:        entity.NewId(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrIDIsInvalid
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price == 0 {
		return ErrPriceIsRequired
	}
	if p.Price < 0 {
		return ErrPriceIsInvalid
	}
	return nil
}