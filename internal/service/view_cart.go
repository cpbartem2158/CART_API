package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/cpbartem2158/CART_API/internal/entity"
	"github.com/cpbartem2158/CART_API/internal/errorsx"
)

func (s *Service) GetCart(ctx context.Context, cartID int64) (*entity.Cart, error) {

	cart, err := s.repo.GetCart(ctx, cartID)
	if err != nil {
		if errors.Is(err, errorsx.ErrCartNotFound) {
			return nil, err
		}
		s.logger.Error("failed to view cart", "error", err)
		return nil, err
	}
	s.logger.Info(fmt.Sprintf("view cart with id: %d", cart.ID))
	return cart, nil
}
