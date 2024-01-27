package model

import "time"

type Order struct {
	ID          string    `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Status      string    `json:"status"`
	CustomerID  string    `json:"customer_id"`
	TotalAmount float64   `json:"total_amount"`
	Note        string    `json:"note"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	OrderItems  []OrderItem
}

type OrderItem struct {
	ID        string    `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	OrderId   string    `json:"order_id"`
	ProductId string    `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	Total     float64   `json:"total"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (OrderItem) TableName() string {
	return "order_item"
}
