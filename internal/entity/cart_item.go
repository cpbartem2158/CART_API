package entity

import "time"

type CartItem struct {
	ID        int       `json:"id" db:"id"`
	CartID    int       `json:"cart_id" db:"cart_id"`
	Product   string    `json:"product" db:"product"`
	Price     float64   `json:"price" db:"price"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
