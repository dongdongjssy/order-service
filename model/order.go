package model

// define order structure
type Order struct {
	CustomerId string `json:"customerId"`
	OrderId    string `json:"orderId"`
	Timestamp  string `json:"timestamp"`
	Items      []Item `json:"items"`
}

// use in-memory array to store all orders
var orders = []Order{}

func (order *Order) Save() error {
	orders = append(orders, *order)
	return nil
}

func GetOrders() []Order {
	return orders
}
