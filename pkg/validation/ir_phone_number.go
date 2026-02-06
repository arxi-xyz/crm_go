package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var iranMobileRe = regexp.MustCompile(`^09\d{9}$`)

func IrPhoneNumberRule(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	return iranMobileRe.MatchString(phone)
}
