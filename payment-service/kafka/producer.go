package kafka

import (
	"context"
	"encoding/json"
	"net"
	"payment-service/pkg/constant"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type Producer struct {
	addr     net.Addr
	balancer kafka.Balancer
}

type IProducer interface {
	PublishOrderRejectTopic(ctx context.Context, orderId string, note string)
	PublishOrderPaidTopic(ctx context.Context, orderId string, data interface{})
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

func (p Producer) PublishOrderRejectTopic(ctx context.Context, orderId string, note string) {
	//publish mess to rollback, rejected order
	data := OrderRejectedMessage{
		OrderId: orderId,
		Note:    note,
	}

	err := p.publish(ctx, constant.OrderRejectedTopic, orderId, data)
	if err != nil {
		logrus.Error("Publish error: " + err.Error())
	}
}
func (p Producer) PublishOrderPaidTopic(ctx context.Context, orderId string, data interface{}) {
	//Update status for order and then update inventory
	_ = p.publish(ctx, constant.OrderPaidTopic, orderId, data)
}
