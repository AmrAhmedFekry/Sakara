package Sellers

import (
	"GoGin/Models"
)

func ProductResource(product Models.Product) map[string]interface{} {
	productResource := make(map[string]interface{})
	productResource["id"] = product.ID
	productResource["cost"] = product.Cost
	productResource["product_name"] = product.ProductName
	productResource["seller_id"] = product.SellerID
	productResource["amount_available"] = product.AmountAvailable

	return productResource
}

func ProductsResource(products []Models.Product) []map[string]interface{} {
	mappedProducts := make([]map[string]interface{}, 0)
	for _, product := range products {
		mappedProducts = append(mappedProducts, ProductResource(product))
	}
	return mappedProducts
}
