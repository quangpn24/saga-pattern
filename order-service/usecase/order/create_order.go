package order

import (
	"context"
	"errors"
	kafka2 "order-service/kafka"
	"order-service/model"
	"order-service/payload"
	"order-service/pkg/constant"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
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

	//publish mess -> payment service
	data := kafka2.OrderCreatedMessage{
		OrderId:     newOrder.ID,
		CustomerId:  newOrder.CustomerID,
		TotalAmount: newOrder.TotalAmount,
	}
	for _, item := range newOrder.OrderItems {
		data.Items = append(data.Items, kafka2.OrderItemMessage{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		})
	}

	var err error
	const retries = 5
	for i := 0; i < retries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// attempt to create topic prior to publishing the message
		err = uc.producer.Publish(ctx, constant.OrderCreatedTopic, newOrder.ID, data)
		if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
			time.Sleep(time.Millisecond * 250)
			continue
		}

		if err != nil {
			logrus.Fatalf("unexpected error %v", err)
		}
		break
	}
	return nil
}
