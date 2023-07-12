package entity

import (
	"testing"

	internalErrors "github.com/mathsant/golang-api/internal/errors"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Panela", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "Panela", product.Name)
	assert.Equal(t, 10.0, product.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 10)
	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, internalErrors.ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Canela", 0)
	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, internalErrors.ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("Canela", -10)
	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, internalErrors.ErrPriceIsInvalid, err)
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct("Product Test", 100)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
