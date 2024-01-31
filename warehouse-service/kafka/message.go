package kafka

type OrderRejectedMessage struct {
	OrderId string `json:"order_id"`
	Note    string `json:"note"`
}

type OrderItemMessage struct {
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
type OrderCreatedMessage struct {
	OrderId     string             `json:"order_id"`
	CustomerId  string             `json:"customer_id"`
	TotalAmount float64            `json:"total_amount"`
	Items       []OrderItemMessage `json:"items"`
}

type OrderPaidMessage struct {
	OrderId       string             `json:"order_id"`
	TransactionId string             `json:"transaction_id"`
	CustomerId    string             `json:"customer_id"`
	Items         []OrderItemMessage `json:"items"`
}
type RefundMessage struct {
	OrderId       string `json:"order_id"`
	TransactionId string `json:"transaction_id"`
	Note          string `json:"note"`
}

type PreparedMessage struct {
	OrderId    string `json:"order_id"`
	CustomerId string `json:"customer_id"`
}
