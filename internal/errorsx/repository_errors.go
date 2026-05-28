package errorsx

import "errors"

var (
	ErrCartNotFound     = errors.New("cart not found")
	ErrCartItemNotFound = errors.New("cart item not found")
	ErrCartFull         = errors.New("cart is full")
)
