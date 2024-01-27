package model

import "time"

type TransType string

const (
	TRANS_IN  TransType = "in"
	TRANS_OUT TransType = "out"
)

type Transaction struct {
	Id         string    `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	CustomerId string    `json:"customer_id"`
	OrderId    string    `json:"order_id"`
	Amount     float64   `json:"amount"`
	TransType  TransType `json:"trans_type"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
