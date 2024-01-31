package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"warehouse-service/config"
	kafka2 "warehouse-service/kafka"
	"warehouse-service/model"
	"warehouse-service/pkg/constant"
	"warehouse-service/usecase"

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
			GroupTopics:    []string{constant.OrderPaidTopic},
			GroupID:        "warehouse-kafka-group",
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
		ctx := context.Background()
		producer := kafka2.NewProducer()
		switch m.Topic {
		case constant.OrderPaidTopic:
			var req kafka2.OrderPaidMessage
			_ = json.Unmarshal(m.Value, &req)

			products := make([]model.Product, 0)
			for _, item := range req.Items {
				products = append(products, model.Product{
					Id:       item.ProductId,
					Quantity: item.Quantity,
				})
			}

			err := c.uc.ProductUC.UpdateQuantity(ctx, products)
			if err != nil {
				//refund
				producer.PublishRefundTopic(ctx, req.OrderId, req.TransactionId, "Product preparation process failed")
			} else {
				producer.PublishOrderPreparedTopic(ctx, req.OrderId, req.CustomerId)
			}

		default:
			fmt.Printf("message at offset %d: %s = %s. Topic: %v\n", m.Offset, string(m.Key), string(m.Value), m.Topic)
		}
	}
}
