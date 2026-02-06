package validation

import "github.com/go-playground/validator/v10"

var v *validator.Validate

func Init() *validator.Validate {
	if v == nil {
		v = validator.New()
		RegisterRules(v)
	}
	return v
}

func V() *validator.Validate {
	if v == nil {
		panic("validation not initialized: call validation.Init()")
	}
	return v
}

func RegisterRules(v *validator.Validate) {
	rules := map[string]validator.Func{
		"ir_phone_number": IrPhoneNumberRule,
	}

	for tag, fn := range rules {
		_ = v.RegisterValidation(tag, fn)
	}
}
