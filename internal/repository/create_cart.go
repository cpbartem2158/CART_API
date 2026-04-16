package repository

import (
	"context"
	"time"

	"github.com/cpbartem2158/CART_API/internal/entity"
)

func (r *Repository) AddCart(ctx context.Context) (*entity.Cart, error) {

	query := `
			INSERT INTO carts (created_at, updated_at) VALUES ($1, $2)
			RETURNING id, created_at, updated_at`

	now := time.Now()

	var cart entity.Cart

	err := r.db.QueryRowContext(ctx, query, now, now).Scan(
		&cart.ID,
		&cart.CreatedAt,
		&cart.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	cart.Items = []entity.CartItem{}
	return &cart, nil
}
