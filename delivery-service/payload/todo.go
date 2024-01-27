package payload

import "github.com/go-playground/validator/v10"

type CreateTodoRequest struct {
	Content string `json:"content" validate:"required"`
	Note    string `json:"note"`
}

func (r *CreateTodoRequest) Validate(validate *validator.Validate) error {
	if err := validate.Struct(r); err != nil {
		//can custom error here
		return err
	}
	return nil
}
