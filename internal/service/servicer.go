package service

import (
	"context"

	"github.com/cpbartem2158/CART_API/internal/entity"
)

type Servicer interface {
	CreateCart(ctx context.Context) (*entity.Cart, error)
	AddCartItemToCart(ctx context.Context, cartID int64, product string, price float64) (*entity.CartItem, error)
	RemoveItem(ctx context.Context, cartID int64, cartItemID int64) error
	GetCart(ctx context.Context, cartID int64) (*entity.Cart, error)
	CalculatePrice(ctx context.Context, cartID int64) (*entity.PriceResponse, error)
}
