package service

import (
	"context"
	"errors"

	"github.com/cpbartem2158/CART_API/internal/entity"
	"github.com/cpbartem2158/CART_API/internal/errorsx"
)

func (s *Service) CalculatePrice(ctx context.Context, cartID int) (*entity.PriceResponse, error) {

	cart, err := s.repo.GetCart(ctx, cartID)
	if err != nil {
		if errors.Is(err, errorsx.ErrCartNotFound) {
			s.logger.Warn("cart not found", "cart_id", cartID)
		} else {
			s.logger.Error("failed to get cart for calculate price", "error", err)
		}
		return nil, err
	}
	var totalPrice float64

	for _, item := range cart.Items {
		totalPrice += item.Price
	}

	discountPercent := 0

	if len(cart.Items) >= 3 && totalPrice > 5000 {
		discountPercent = 15
	} else if totalPrice > 5000 {
		discountPercent = 10
	} else if len(cart.Items) >= 3 {
		discountPercent = 5
	}

	finalPrice := totalPrice * (1 - float64(discountPercent)/100)

	response := &entity.PriceResponse{
		CartID:          cartID,
		TotalPrice:      totalPrice,
		DiscountPercent: discountPercent,
		FinalPrice:      finalPrice,
	}
	return response, nil

}
