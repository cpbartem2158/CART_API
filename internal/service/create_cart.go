package service

import (
	"context"
	"fmt"

	"github.com/cpbartem2158/CART_API/internal/entity"
)

func (s *Service) CreateCart(ctx context.Context) (*entity.Cart, error) {

	cart, err := s.repo.AddCart(ctx)
	if err != nil {
		s.logger.Error("failed to create cart", err)
		return nil, err
	}
	s.logger.Info(fmt.Sprintf("created cart with id: %d", cart.ID))
	return cart, nil

}
