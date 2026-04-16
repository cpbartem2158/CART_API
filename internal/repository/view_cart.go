package repository

import (
	"context"
	"database/sql"

	"github.com/cpbartem2158/CART_API/internal/entity"
	"github.com/cpbartem2158/CART_API/internal/errorsx"
)

func (r *Repository) GetCart(ctx context.Context, cartID int) (*entity.Cart, error) {

	transaction, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer transaction.Rollback()

	var exists bool
	checkCartExistQuery := `SELECT EXISTS (SELECT 1 FROM carts WHERE id = $1)`
	err = transaction.QueryRowContext(ctx, checkCartExistQuery, cartID).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errorsx.ErrCartNotFound
	}

	viewCartQuery := `
					SELECT id, created_at, updated_at FROM carts WHERE id = $1`

	var cart entity.Cart

	err = transaction.QueryRowContext(ctx, viewCartQuery, cartID).Scan(
		&cart.ID,
		&cart.CreatedAt,
		&cart.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errorsx.ErrCartNotFound
	}
	if err != nil {
		return nil, err
	}
	cartItemsQuery := `SELECT id , cart_id, product, price, created_at, updated_at FROM cart_items WHERE cart_id = $1 ORDER BY id`

	rows, err := transaction.QueryContext(ctx, cartItemsQuery, cartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cartItems []entity.CartItem
	for rows.Next() {
		var item entity.CartItem
		err := rows.Scan(
			&item.ID,
			&item.CartID,
			&item.Product,
			&item.Price,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		cartItems = append(cartItems, item)
	}
	cart.Items = cartItems
	if err := transaction.Commit(); err != nil {
		return nil, err
	}

	return &cart, nil
}
