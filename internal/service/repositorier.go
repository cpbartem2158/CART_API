package service

import (
	"context"

	"github.com/cpbartem2158/CART_API/internal/entity"
)

type Repositorier interface {
	AddCart(ctx context.Context) (*entity.Cart, error)
	AddCartItem(ctx context.Context, cartID int, product string, price float64) (*entity.CartItem, error)
}
