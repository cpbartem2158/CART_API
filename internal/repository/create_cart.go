package repository

import (
	"context"

	"github.com/cpbartem2158/CART_API/internal/entity"
)

func (r *Repository) CreateCart(ctx context.Context) (*entity.Cart, error) {

	query := `
   INSERT INTO carts DEFAULT VALUES
   RETURNING id, created_at, updated_at
`
	var cart entity.Cart

	err := r.db.QueryRowContext(ctx, query).Scan(
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
