package service

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/cpbartem2158/CART_API/internal/entity"
	"github.com/cpbartem2158/CART_API/internal/errorsx"
	"github.com/cpbartem2158/CART_API/internal/service/mocks"
	"github.com/stretchr/testify/assert"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	Level: slog.LevelDebug,
}))

func TestService_CreateCart(t *testing.T) {
	mockRepo := mocks.NewMockRepository(t)

	expectedCart := &entity.Cart{
		ID:    1,
		Items: []entity.CartItem{},
	}

	mockRepo.On("CreateCart", context.Background()).Return(expectedCart, nil)

	svc := NewService(mockRepo, logger)
	cart, err := svc.CreateCart(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, 1, cart.ID)
	assert.Empty(t, cart.Items)

}

func TestService_ViewCart_Success(t *testing.T) {
	mockRepo := mocks.NewMockRepository(t)
	expectedCart := &entity.Cart{
		ID: 1,
		Items: []entity.CartItem{
			{
				ID:      1,
				CartID:  1,
				Product: "test",
				Price:   121,
			},
		},
	}
	mockRepo.On("GetCart", context.Background(), 1).Return(expectedCart, nil)

	svc := NewService(mockRepo, logger)

	cart, err := svc.GetCart(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, expectedCart, cart)

}

func TestService_ViewCart_Fail(t *testing.T) {
	mockRepo := mocks.NewMockRepository(t)

	mockRepo.On("GetCart", context.Background(), 191).Return(nil, errorsx.ErrCartNotFound)

	svc := NewService(mockRepo, logger)
	cart, err := svc.GetCart(context.Background(), 191)

	assert.Error(t, err)
	assert.Nil(t, cart)
	assert.Equal(t, errorsx.ErrCartNotFound, err)

}

func TestService_AddItem_Success(t *testing.T) {
	mockRepo := mocks.NewMockRepository(t)

	expectedItem := &entity.CartItem{
		ID:      1,
		CartID:  1,
		Product: "test",
		Price:   121,
	}
	mockRepo.On("AddCartItem", context.Background(), 1, "test", 121.0).Return(expectedItem, nil)

	svc := NewService(mockRepo, logger)

	item, err := svc.AddCartItemToCart(context.Background(), 1, "test", 121.0)

	assert.NoError(t, err)
	assert.Equal(t, expectedItem, item)

}

func TestService_AddItem_FullCart(t *testing.T) {
	mockRepo := mocks.NewMockRepository(t)

	mockRepo.On("AddCartItem", context.Background(), 1, "test", 121.0).Return(nil, errorsx.ErrCartFull)

	svc := NewService(mockRepo, logger)

	item, err := svc.AddCartItemToCart(context.Background(), 1, "test", 121.0)
	assert.Error(t, err)
	assert.Nil(t, item)
	assert.Equal(t, errorsx.ErrCartFull, err)
}

func TestService_CalculatePrice(t *testing.T) {
	mockRepo := mocks.NewMockRepository(t)

	cart := &entity.Cart{
		ID: 1,
		Items: []entity.CartItem{
			{
				ID:      1,
				Product: "test",
				Price:   100.0,
			},
			{
				ID:      2,
				Product: "test2",
				Price:   5122.0,
			},
		},
	}
	mockRepo.On("GetCart", context.Background(), 1).Return(cart, nil)
	svc := NewService(mockRepo, logger)
	priceInfo, err := svc.CalculatePrice(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, 5222.0, priceInfo.TotalPrice)
	assert.Equal(t, 10, priceInfo.DiscountPercent)
	assert.Equal(t, 5222.0*0.9, priceInfo.FinalPrice)

}

func TestService_RemoveItem_Success(t *testing.T) {
	mockRepo := mocks.NewMockRepository(t)

	mockRepo.On("RemoveCartItem", context.Background(), 1, 1).Return(nil)

	svc := NewService(mockRepo, logger)

	err := svc.RemoveItem(context.Background(), 1, 1)
	assert.NoError(t, err)
}

func TestService_RemoveItem_CartNotFound(t *testing.T) {
	mockRepo := mocks.NewMockRepository(t)

	mockRepo.On("RemoveCartItem", context.Background(), 1, 1).Return(errorsx.ErrCartNotFound)

	svc := NewService(mockRepo, logger)

	err := svc.RemoveItem(context.Background(), 1, 1)
	assert.Error(t, err)
}

func TestService_RemoveItem_ItemNotFound(t *testing.T) {
	mockRepo := mocks.NewMockRepository(t)

	mockRepo.On("RemoveCartItem", context.Background(), 1, 1).Return(errorsx.ErrCartItemNotFound)
	svc := NewService(mockRepo, logger)
	err := svc.RemoveItem(context.Background(), 1, 1)
	assert.Error(t, err)
	
}
