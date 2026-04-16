package service

import (
	"context"

	"github.com/cpbartem2158/CART_API/internal/entity"
)

type Repositorier interface {
	AddCart(ctx context.Context) (*entity.Cart, error)
}
