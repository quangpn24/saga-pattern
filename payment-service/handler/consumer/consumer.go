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
			GroupTopics:    []string{constant.OrderCreatedTopic, constant.RefundTopic},
			GroupID:        "payment-kafka-group",
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
		producer := kafka2.NewProducer()

		switch m.Topic {
		case constant.OrderCreatedTopic:
			//payment
			var req kafka2.OrderCreatedMessage
			_ = json.Unmarshal(m.Value, &req)
			transId, err := c.uc.CustomerUseCase.PayTheBill(ctx, req)

			if err != nil {
				logrus.Error("Error during payment")

				producer.PublishOrderRejectTopic(ctx, req.OrderId, "Payment fail")
			} else {
				data := kafka2.OrderPaidMessage{
					OrderId:       req.OrderId,
					CustomerId:    req.CustomerId,
					TransactionId: transId,
					Items:         req.Items,
				}
				producer.PublishOrderPaidTopic(ctx, req.OrderId, data)
			}
		case constant.RefundTopic:
			var req kafka2.RefundMessage
			_ = json.Unmarshal(m.Value, &req)
			_ = c.uc.CustomerUseCase.Refund(ctx, req.TransactionId)
			producer.PublishOrderRejectTopic(ctx, req.OrderId, req.Note)
		default:
			fmt.Printf("message at offset %d: %s = %s. Topic: %v\n", m.Offset, string(m.Key), string(m.Value), m.Topic)
		}
	}
}
