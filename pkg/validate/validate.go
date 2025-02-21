package validate

import (
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func SetupValidate(v *validator.Validate) {
	Validate = v
	Validate.RegisterValidation("test", testValidation)
}

func testValidation(fl validator.FieldLevel) bool {
	return true
}
