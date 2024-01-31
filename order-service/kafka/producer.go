package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"net"
	"order-service/model"
	"order-service/pkg/constant"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type Producer struct {
	addr     net.Addr
	balancer kafka.Balancer
}

type IProducer interface {
	PublishOrderCreatedTopic(order model.Order)
}

func NewProducer() IProducer {
	return &Producer{
		addr:     kafka.TCP("127.0.0.1:9092", "127.0.0.1:9093", "127.0.0.1:9094"),
		balancer: &kafka.Hash{},
	}
}
func (p Producer) publish(ctx context.Context, topic string, key string, data interface{}) error {
	w := &kafka.Writer{
		Addr:     p.addr,
		Topic:    topic,
		Balancer: p.balancer,
	}

	mess, _ := json.Marshal(data)
	err := w.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte(key),
			Value: mess,
		},
	)
	if err != nil {
		return err
	}

	if err := w.Close(); err != nil {
		logrus.Fatal("failed to close writer:", err)
	}
	return nil
}

func (p Producer) PublishOrderCreatedTopic(order model.Order) {
	//publish mess -> payment service
	data := OrderCreatedMessage{
		OrderId:     order.ID,
		CustomerId:  order.CustomerID,
		TotalAmount: order.TotalAmount,
	}
	for _, item := range order.OrderItems {
		data.Items = append(data.Items, OrderItemMessage{
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
		err = p.publish(ctx, constant.OrderCreatedTopic, order.ID, data)
		if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
			time.Sleep(time.Millisecond * 250)
			continue
		}

		if err != nil {
			logrus.Fatalf("unexpected error %v", err)
		}
		break
	}
}
