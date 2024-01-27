package order

import (
	"context"
	"errors"
	"net/http"
	"order-service/pkg/appmodel"
	"order-service/pkg/constant"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (uc *UseCase) CancelOrder(ctx context.Context, id string) error {
	order, err := uc.GetById(ctx, id)
	if err != nil && err != gorm.ErrRecordNotFound {
		return appmodel.AppError{Code: http.StatusInternalServerError, Message: err.Error(), Err: err}
	}
	if err == gorm.ErrRecordNotFound {
		logrus.Error(err.Error())
		return appmodel.AppError{Code: http.StatusBadRequest, Message: err.Error(), Err: err}
	}

	//check order status
	if order.Status != string(constant.ORDER_COMPLETED) {
		err := errors.New("the order is still in the process of being created")
		logrus.Error(err.Error())
		return appmodel.AppError{Code: http.StatusBadRequest, Message: err.Error(), Err: err}
	}

	order.Status = string(constant.ORDER_CANCEL)
	if err := uc.repo.UpdateOrder(ctx, order); err != nil {
		logrus.Error(err.Error())
		return appmodel.AppError{Code: http.StatusInternalServerError, Message: err.Error(), Err: err}
	}
	return nil
}
