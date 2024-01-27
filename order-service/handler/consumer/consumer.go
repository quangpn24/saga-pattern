package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"order-service/config"
	kafka2 "order-service/kafka"
	"order-service/pkg/constant"
	"order-service/usecase"
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
			GroupTopics:    []string{constant.OrderRejectedTopic, constant.OrderPaidTopic},
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
		m, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			logrus.Fatal("err: " + err.Error())
			break
		}
		switch m.Topic {
		case constant.OrderRejectedTopic:
			var req kafka2.OrderRejectedMessage
			_ = json.Unmarshal(m.Value, &req)
			if err := c.uc.OrderUseCase.RejectOrder(context.Background(), req.OrderId, req.Note); err != nil {
				logrus.Error("reject order fail: " + err.Error())
			}
		case constant.OrderPaidTopic:
			var req kafka2.OrderDoneStep
			_ = json.Unmarshal(m.Value, &req)
			if err := c.uc.OrderUseCase.UpdateStatus(context.Background(), req.OrderId, req.Status); err != nil {
				logrus.Error("update order fail: " + err.Error())
			}
		default:
			fmt.Printf("message at offset %d: %s = %s. Topic: %v\n", m.Offset, string(m.Key), string(m.Value), m.Topic)
		}
	}
}
