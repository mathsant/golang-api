package errors

import "errors"

var (
	ErrIDIsRequired    = errors.New("id is required")
	ErrIDIsInvalid     = errors.New("invalid id")
	ErrNameIsRequired  = errors.New("name is requrired")
	ErrPriceIsRequired = errors.New("price is requrired")
	ErrPriceIsInvalid  = errors.New("invalid price")
)
