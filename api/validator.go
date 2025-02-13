package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/sowmiyaramanathan/A-simple-bank-app-using-Golang/db/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
