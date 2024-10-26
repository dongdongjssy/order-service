package model

type Summary struct {
	CustomerId          string  `json:"customerId"`
	NbrOfPurchasedItems uint    `json:"nbrOfPurchasedItems"`
	TotalAmountEur      float64 `json:"totalAmountEur"`
}
