package repository

import (
	"context"
	"payment-service/model"
	"payment-service/pkg/constant"

	"gorm.io/gorm"
)

func (r *Repository) PayTheBill(ctx context.Context, trans model.Transaction) error {
	tx := r.db.WithContext(ctx)

	tx = tx.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	//update balance for customer
	if err := tx.Table(constant.CustomerTable).Where("id = ?", trans.CustomerId).Updates(map[string]interface{}{"balance": gorm.Expr("balance - ?", trans.Amount)}).Error; err != nil {
		tx.Rollback()
		return err
	}

	//create new transaction
	if err := tx.Table(constant.TransactionTable).Create(&trans).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *Repository) GetCustomerById(ctx context.Context, id string) (*model.Customer, error) {
	var customer *model.Customer
	err := r.db.WithContext(ctx).Table(constant.CustomerTable).Where("id = ?", id).Take(&customer).Error
	return customer, err
}
func (r *Repository) Refund(ctx context.Context, transId string) error {
	tx := r.db.WithContext(ctx)
	tx = tx.Begin()

	//get amount
	var trans model.Transaction
	if err := tx.Table(constant.TransactionTable).Where("id = ?", transId).Select("*").Take(&trans).Error; err != nil {
		tx.Rollback()
		return err
	}

	//delete transaction
	if err := tx.Table(constant.TransactionTable).Where("id = ?", transId).Delete(&model.Transaction{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	//refund
	if err := tx.Table(constant.CustomerTable).Where("id = ?", trans.CustomerId).
		Updates(map[string]interface{}{"balance": gorm.Expr("balance + ?", trans.Amount)}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
