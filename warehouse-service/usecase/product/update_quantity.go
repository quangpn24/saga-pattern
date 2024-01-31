package product

import (
	"context"
	"errors"
	"time"
	"warehouse-service/model"
)

func (uc *UseCase) UpdateQuantity(ctx context.Context, products []model.Product) error {
	var (
		err     error
		product *model.Product
	)
	time.Sleep(time.Minute)
	if len(products) == 0 {
		return nil
	}
	for i, p := range products {
		product, err = uc.repo.GetProductById(ctx, p.Id)
		if err != nil {
			return err
		}

		if product.Quantity < p.Quantity {
			err := errors.New("not enough quantity in stock")
			return err
		}

		product.Quantity -= p.Quantity
		products[i] = *product
	}

	err = uc.repo.UpdateProducts(ctx, products)
	if err != nil {
		return err
	}
	return nil
}
