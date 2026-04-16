package service

import (
	"context"

	"github.com/cpbartem2158/CART_API/internal/entity"
)

type Servicer interface {
	CreateCart(ctx context.Context) (*entity.Cart, error)
}
