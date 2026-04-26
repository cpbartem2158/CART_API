package repository

import (
	"context"
	"time"

	"github.com/cpbartem2158/CART_API/internal/entity"
	"github.com/cpbartem2158/CART_API/internal/errorsx"
)

func (r *Repository) GetCart(ctx context.Context, cartID int) (*entity.Cart, error) {

	query := `
   SELECT 
       c.id, c.created_at, c.updated_at,
       ci.id, ci.cart_id, ci.product, ci.price, ci.created_at,ci.updated_at
   FROM carts c
   LEFT JOIN cart_items ci ON c.id = ci.cart_id
   WHERE c.id = $1`

	rows, err := r.db.QueryContext(ctx, query, cartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cart := &entity.Cart{}
	var items []entity.CartItem
	cartFound := false

	for rows.Next() {
		var itemID, itemCartID *int64
		var product *string
		var price *float64
		var itemCreatedAt, itemUpdatedAt *time.Time

		err := rows.Scan(
			&cart.ID,
			&cart.CreatedAt,
			&cart.UpdatedAt,
			&itemID,
			&itemCartID,
			&product,
			&price,
			&itemCreatedAt,
			&itemUpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		cartFound = true

		if itemID != nil {
			items = append(items, entity.CartItem{
				ID:        *itemID,
				CartID:    *itemCartID,
				Product:   *product,
				Price:     *price,
				CreatedAt: *itemCreatedAt,
				UpdatedAt: *itemUpdatedAt,
			})
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if !cartFound {
		return nil, errorsx.ErrCartNotFound
	}
	if items == nil {
		items = []entity.CartItem{}
	}
	cart.Items = items
	return cart, nil
}
