package domain

type Order struct {
	ID         int    `json:"id"`
	CustomerID int    `json:"customer_id"`
	OrderDate  string `json:"order_date"`
}
