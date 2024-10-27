package model

type Summary struct {
	CustomerId          string  `json:"customerId" binding:"required"`
	NbrOfPurchasedItems int     `json:"nbrOfPurchasedItems" binding:"required"`
	TotalAmountEur      float64 `json:"totalAmountEur" binding:"required"`
	Items               []Item  `json:"items" binding:"required,dive"`
}
