package service

import (
	"context"
	"fmt"
)

func (s *Service) RemoveItem(ctx context.Context, cartID int, cartItemID int) error {

	err := s.repo.RemoveCartItem(ctx, cartID, cartItemID)
	if err != nil {
		s.logger.Error("failed remove item", "error", err)
		return err
	}
	s.logger.Info(fmt.Sprintf("remove item with id: %d from cart with id: %d", cartItemID, cartID))
	return nil
}
