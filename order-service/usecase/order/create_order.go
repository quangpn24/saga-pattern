package order

import (
	"context"
	"order-service/model"
	"order-service/payload"
	"order-service/pkg/constant"

	"github.com/google/uuid"
)

func (uc *UseCase) CreateOrder(ctx context.Context, req payload.CreateOrderRequest) error {
	orderId := uuid.New().String()
	newOrder := &model.Order{
		ID:         orderId,
		Status:     constant.ORDER_CREATED,
		CustomerID: "d05fcbfd-4c18-4242-8601-925af87504e0",
	}

	totalAmount := 0.0
	for _, item := range req.Items {
		totalAmount += item.Total
		newOrder.OrderItems = append(newOrder.OrderItems, model.OrderItem{
			OrderId:   orderId,
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
			Price:     item.Price,
			Total:     item.Total,
		})
	}
	newOrder.TotalAmount = totalAmount

	if err := uc.repo.CreateOrder(ctx, newOrder); err != nil {
		return err
	}

	//publish message
	uc.producer.PublishOrderCreatedTopic(*newOrder)
	return nil
}
