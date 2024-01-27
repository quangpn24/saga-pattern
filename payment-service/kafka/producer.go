package kafka

import (
	"context"
	"encoding/json"
	"net"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type Producer struct {
	addr     net.Addr
	balancer kafka.Balancer
}

type IProducer interface {
	Publish(ctx context.Context, topic string, key string, data interface{}) error
}

func NewProducer() IProducer {
	return &Producer{
		addr:     kafka.TCP("127.0.0.1:9092", "127.0.0.1:9093", "127.0.0.1:9094"),
		balancer: &kafka.Hash{},
	}
}
func (p Producer) Publish(ctx context.Context, topic string, key string, data interface{}) error {
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
