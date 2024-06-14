package customvalidator

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func CodeNumberId(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()

	return strings.EqualFold(phoneNumber[:2], "62")
}
