package repository

import (
	"context"

	"github.com/cpbartem2158/CartAPI/internal/errorsx"
)

func (r *Repository) RemoveCartItem(ctx context.Context, cartID int, cartItemID int) error {

	transaction, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer transaction.Rollback()

	var exists bool
	checkCartExistQuery := `SELECT EXISTS (SELECT 1 FROM carts WHERE id = $1)`
	err = transaction.QueryRowContext(ctx, checkCartExistQuery, cartID).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return errorsx.ErrCartNotFound
	}

	removeItemFromCartQuery := `
					DELETE FROM cart_items WHERE id = $1 AND cart_id = $2`

	result, err := transaction.ExecContext(ctx, removeItemFromCartQuery, cartItemID, cartID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errorsx.ErrCartItemNotFound
	}

	if err := transaction.Commit(); err != nil {
		return err
	}

	return nil

}
