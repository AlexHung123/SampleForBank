package api

import (
	"github.com/alexhung/simplebank/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		//check currency support
		return util.IsSupportedCurrency(currency)
	}
	return false
}
