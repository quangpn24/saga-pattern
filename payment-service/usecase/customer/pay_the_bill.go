package customer

import (
	"context"
	"fmt"
	kafka2 "payment-service/kafka"
	"payment-service/model"

	"github.com/sirupsen/logrus"
)

func (uc *UseCase) PayTheBill(ctx context.Context, req kafka2.OrderCreatedMessage) error {
	//get customer balance
	customer, err := uc.repo.GetCustomerById(ctx, req.CustomerId)
	if err != nil {
		logrus.Error("Error when getting customer by id: " + err.Error())
		return err
	}

	if customer.Balance < req.TotalAmount {
		err := fmt.Errorf("account has insufficient funds")
		logrus.Error(err.Error())
		return err
	}
	newTransaction := model.Transaction{
		CustomerId: req.CustomerId,
		OrderId:    req.OrderId,
		Amount:     req.TotalAmount,
		TransType:  model.TRANS_OUT,
		Content:    "Pay the bill",
	}
	return uc.repo.PayTheBill(ctx, newTransaction)
}
