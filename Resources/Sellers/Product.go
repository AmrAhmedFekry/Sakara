package Sellers

import (
	"GoGin/Models"
)

/*
	The Idea behind this snippet is to use this function to return only required fields
	of every model in case of single product, list of products and depend on the API requested Actor
*/

// Map of  product data in case of single product and if there are common fields
// Between singel and list of products then use this function to set them
func ProductResource(product Models.Product) map[string]interface{} {
	productResource := make(map[string]interface{})
	productResource["id"] = product.ID
	productResource["cost"] = product.Cost
	productResource["product_name"] = product.ProductName
	productResource["seller_id"] = product.SellerID
	productResource["amount_available"] = product.AmountAvailable

	return productResource
}

// Map of products data
func ProductsResource(products []Models.Product) []map[string]interface{} {
	mappedProducts := make([]map[string]interface{}, 0)
	for _, product := range products {
		mappedProducts = append(mappedProducts, ProductResource(product))
	}
	return mappedProducts
}
