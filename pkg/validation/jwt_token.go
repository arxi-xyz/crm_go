package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// todo: maybe you can implement infra for this regex checks
var jwtTokenValidation = regexp.MustCompile(`^[A-Za-z0-9_-]+\.[A-Za-z0-9_-]+\.[A-Za-z0-9_-]+$`)

func JwtTokenRule(fl validator.FieldLevel) bool {
	token := fl.Field().String()
	return jwtTokenValidation.MatchString(token)
}
