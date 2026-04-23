package service

import (
	"context"

	"github.com/cpbartem2158/CART_API/internal/entity"
)

type Servicer interface {
	CreateCart(ctx context.Context) (*entity.Cart, error)
	AddCartItemToCart(ctx context.Context, cartID int, product string, price float64) (*entity.CartItem, error)
	RemoveItem(ctx context.Context, cartID int, cartItemID int) error
	GetCart(ctx context.Context, cartID int) (*entity.Cart, error)
}
