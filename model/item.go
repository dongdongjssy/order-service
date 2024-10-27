package model

type Item struct {
	CustomerId string  `json:"customerId,omitempty"`
	ItemId     string  `json:"itemId" binding:"required"`
	CostEur    float64 `json:"costEur" binding:"required"`
}
