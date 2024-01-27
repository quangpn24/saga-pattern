package payload

import "github.com/go-playground/validator/v10"

type ItemRequest struct {
	ProductId string  `json:"product_id" validate:"required"`
	Quantity  int     `json:"quantity" validate:"required"`
	Price     float64 `json:"price" validate:"required"`
	Total     float64 `json:"total" validate:"required"`
}
type CreateOrderRequest struct {
	Items []ItemRequest `json:"items" validate:"dive,required"`
}

func (r *CreateOrderRequest) Validate(validate *validator.Validate) error {
	return validate.Struct(r)
}
