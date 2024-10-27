package model

type Order struct {
	CustomerId string `json:"customerId" binding:"required"`
	OrderId    string `json:"orderId" binding:"required"`
	Timestamp  string `json:"timestamp" binding:"required"`
	Items      []Item `json:"items" binding:"required,min=1,dive"`
}
