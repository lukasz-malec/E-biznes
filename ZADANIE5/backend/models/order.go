package models
import "time"

type Customer struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type OrderItem struct {
	ProductID int     `json:"product_id"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type Order struct {
	ID        int         `json:"id"`
	Items     []OrderItem `json:"items"`
	Total     float64     `json:"total"`
	Customer  Customer    `json:"customer"`
	Status    string      `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
}