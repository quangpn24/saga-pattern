package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"payment-service/config"
	kafka2 "payment-service/kafka"
	"payment-service/pkg/constant"
	"payment-service/usecase"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type Consumer struct {
	reader *kafka.Reader
	uc     *usecase.UseCase
	cfg    *config.Config
}

type IConsumer interface {
	Consume()
}

func NewConsumer(uc *usecase.UseCase, cfg *config.Config) IConsumer {
	return &Consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:        []string{"localhost:9092", "localhost:9093", "localhost:9094"},
			GroupTopics:    []string{constant.OrderCreatedTopic},
			GroupID:        "order-kafka-group",
			MaxBytes:       10e6, // 10MB
			CommitInterval: time.Second,
		}),
		uc:  uc,
		cfg: cfg,
	}
}

func (c *Consumer) Consume() {
	for {
		ctx := context.Background()
		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			logrus.Fatal("err: " + err.Error())
			break
		}
		switch m.Topic {
		case constant.OrderCreatedTopic:
			//payment
			var req kafka2.OrderCreatedMessage
			_ = json.Unmarshal(m.Value, &req)
			err := c.uc.CustomerUseCase.PayTheBill(ctx, req)
			if err != nil {
				logrus.Error("Error during payment")

				//publish mess to rollback, rejected order
				data := kafka2.OrderRejectedMessage{
					OrderId: req.OrderId,
					Note:    "Error during payment",
				}
				producer := kafka2.NewProducer()
				err := producer.Publish(ctx, constant.OrderRejectedTopic, req.OrderId, data)
				if err != nil {
					logrus.Error("Publish error: " + err.Error())
				}
			} else {
				//Update status for order and then update inventory
				data := kafka2.OrderDoneStep{
					OrderId: req.OrderId,
					Status:  "ORDER_PAID",
				}
				producer := kafka2.NewProducer()
				_ = producer.Publish(ctx, constant.OrderPaidTopic, req.OrderId, data)
			}
		default:
			fmt.Printf("message at offset %d: %s = %s. Topic: %v\n", m.Offset, string(m.Key), string(m.Value), m.Topic)
		}
	}
}
