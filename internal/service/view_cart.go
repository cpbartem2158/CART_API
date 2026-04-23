package service

import (
	"context"
	"fmt"

	"github.com/cpbartem2158/CART_API/internal/entity"
)

func (s *Service) GetCart(ctx context.Context, cartID int) (*entity.Cart, error) {

	cart, err := s.repo.GetCart(ctx, cartID)
	if err != nil {
		s.logger.Error("failed to view cart", "error", err)
		return nil, err
	}
	s.logger.Info(fmt.Sprintf("view cart with id: %d", cart.ID))
	return cart, nil
}
