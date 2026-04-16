package entity

import "time"

type Cart struct {
	ID        int        `json:"id" db:"id"`
	Items     []CartItem `json:"items" db:"-"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
}
