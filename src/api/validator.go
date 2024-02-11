package api

import (
	"github.com/go-playground/validator/v10"
)

const (
	USD = "USD"
	EUR = "RUE"
	CNY = "CNT"
	GBP = "GBP"
)

func isSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CNY, GBP:
		return true
	}
	return false
}

var validateCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return isSupportedCurrency(currency)
	}
	return false
}
