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
type OrderDoneStep struct {
	OrderId string `json:"order_id"`
	Status  string `json:"status"`
}
