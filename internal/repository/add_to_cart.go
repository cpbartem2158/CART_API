package repository

import (
	"context"

	"github.com/cpbartem2158/CART_API/internal/entity"
	"github.com/cpbartem2158/CART_API/internal/errorsx"
)

func (r *Repository) AddCartItem(ctx context.Context, cartID int, product string, price float64) (*entity.CartItem, error) {

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

	var count int

	countCartItemsQuery := `SELECT COUNT(*) FROM cart_items WHERE cart_id = $1`
	err = transaction.QueryRowContext(ctx, countCartItemsQuery, cartID).Scan(&count)
	if err != nil {
		return nil, err
	}
	if count >= 5 {
		return nil, errorsx.ErrCartFull
	}

	addItemToCartQuery := `
					INSERT INTO cart_items (cart_id, product, price)
					VALUES ($1, $2, $3)
					RETURNING id, cart_id, product, price, created_at, updated_at`

	var item entity.CartItem

	err = transaction.QueryRowContext(ctx, addItemToCartQuery, cartID, product, price).Scan(
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
	if err := transaction.Commit(); err != nil {
		return nil, err
	}
	return &item, nil

}
