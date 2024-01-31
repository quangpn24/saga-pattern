package customer

import (
	"context"
	"fmt"
	kafka2 "payment-service/kafka"
	"payment-service/model"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (uc *UseCase) PayTheBill(ctx context.Context, req kafka2.OrderCreatedMessage) (string, error) {
	//get customer balance
	customer, err := uc.repo.GetCustomerById(ctx, req.CustomerId)
	if err != nil {
		logrus.Error("Error when getting customer by id: " + err.Error())
		return "", err
	}

	if customer.Balance < req.TotalAmount {
		err := fmt.Errorf("account has insufficient funds")
		logrus.Error(err.Error())
		return "", err
	}

	transId := uuid.New().String()
	newTransaction := model.Transaction{
		Id:         transId,
		CustomerId: req.CustomerId,
		OrderId:    req.OrderId,
		Amount:     req.TotalAmount,
		TransType:  model.TRANS_OUT,
		Content:    "Pay the bill",
	}
	err = uc.repo.PayTheBill(ctx, newTransaction)
	return transId, err
}
