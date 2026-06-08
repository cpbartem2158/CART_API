package service

import (
	"context"
	"fmt"

	"github.com/cpbartem2158/CART_API/internal/entity"
)

func (s *Service) CreateCart(ctx context.Context) (*entity.Cart, error) {

	cart, err := s.repo.CreateCart(ctx)
	if err != nil {
		s.logger.Error("failed to create cart", "error", err)
		return nil, err
	}
	s.logger.Info(fmt.Sprintf("created cart with id: %d", cart.ID))
	return cart, nil

}
