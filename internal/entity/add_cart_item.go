package entity

type AddItemRequest struct {
	Product string  `json:"product"`
	Price   float64 `json:"price"`
}

type PriceResponse struct {
	CartID          int     `json:"cart_id"`
	TotalPrice      float64 `json:"total_price"`
	DiscountPercent int     `json:"discount_percent"`
	FinalPrice      float64 `json:"final_price"`
}
