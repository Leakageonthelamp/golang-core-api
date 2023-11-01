package requests

import core "github.com/Leakageonthelamp/go-leakage-core"

type UserCreateRequest struct {
	core.BaseValidator
	FirstName *string `json:"first_name" validate:"required"`
	LastName  *string `json:"last_name" validate:"required"`
	Phone     *string `json:"phone" validate:"required"`
	Address   *string `json:"address" validate:"required"`
}

func (r UserCreateRequest) Valid(ctx core.IContext) core.IError {
	r.Must(r.IsStrRequired(r.FirstName, "first_name"))
	r.Must(r.IsStrRequired(r.LastName, "last_name"))
	r.Must(r.IsStringNumber(r.Phone, "phone"))

	return r.Error()
}
