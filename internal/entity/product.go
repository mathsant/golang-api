package entity

import (
	"time"

	internalErrors "github.com/mathsant/golang-api/internal/errors"
	"github.com/mathsant/golang-api/pkg/entity"
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
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
		return internalErrors.ErrIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return internalErrors.ErrIDIsInvalid
	}
	if p.Name == "" {
		return internalErrors.ErrNameIsRequired
	}
	if p.Price == 0 {
		return internalErrors.ErrPriceIsRequired
	}
	if p.Price < 0 {
		return internalErrors.ErrPriceIsInvalid
	}
	return nil
}
