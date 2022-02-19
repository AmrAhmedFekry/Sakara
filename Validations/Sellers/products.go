package Sellers

import (
	"GoGin/Models"
	"GoGin/Validations"

	validation "github.com/go-ozzo/ozzo-validation"
)

func InsertValidation(product Models.Product) validation.Errors {
	return validation.Errors{
		"amount_available": validation.Validate(product.AmountAvailable, Validations.RequiredRule()),
		"cost":             validation.Validate(product.Cost, Validations.RequiredRule()),
		"product_name":     validation.Validate(product.ProductName, Validations.RequiredRule()),
	}
}
