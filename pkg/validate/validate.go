package validate

import (
	"github.com/go-playground/validator/v10"
)

var V *validator.Validate

func SetupValidate(v *validator.Validate) {
	V = v
	V.RegisterValidation("test", testValidation)
}

func testValidation(fl validator.FieldLevel) bool {
	return true
}
