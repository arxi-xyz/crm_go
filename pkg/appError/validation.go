package appError

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func FromValidator(err error) (*AppError, bool) {
	var verrs validator.ValidationErrors
	if !errors.As(err, &verrs) {
		return nil, false
	}

	fields := map[string][]string{}
	for _, fe := range verrs {
		field := strings.ToLower(fe.Field())
		tag := fe.Tag()
		fields[field] = append(fields[field], tag)
	}

	return Validation("اطلاعات نامعتبر است", fields), true
}
