package service

import (
	"context"
	"fmt"

	"github.com/cpbartem2158/CART_API/internal/entity"
)

func (s *Service) AddCartItemToCart(ctx context.Context, cartID int, product string, price float64) (*entity.CartItem, error) {

	cartItem, err := s.repo.AddCartItem(ctx, cartID, product, price)
	if err != nil {
		s.logger.Error("failed to add item to cart", err)
		return nil, err
	}
	s.logger.Info(fmt.Sprintf("add item with id: %d", cartItem.ID))
	return cartItem, nil
}
