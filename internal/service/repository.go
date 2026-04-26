package service

import (
	"context"

	"github.com/cpbartem2158/CART_API/internal/entity"
)

//go:generate mockery --name=Repository --output=mocks --outpkg=mocks --filename=mock_repository.go --with-expecter
type Repository interface {
	CreateCart(ctx context.Context) (*entity.Cart, error)
	AddCartItem(ctx context.Context, cartID int, product string, price float64) (*entity.CartItem, error)
	RemoveCartItem(ctx context.Context, cartID int, cartItemID int) error
	GetCart(ctx context.Context, cartID int) (*entity.Cart, error)
}
