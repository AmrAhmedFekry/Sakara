package Visitors

import (
	"GoGin/Application"
	"GoGin/Models"
	Resources "GoGin/Resources/Buyers"

	"github.com/gin-gonic/gin"
)

// Show list of products
func ListOfProducts(c *gin.Context) {
	r := Application.NewRequest(c)

	var products []Models.Product
	r.DB.Find(&products)
	// Use Of resource to return only required fields to show
	r.Success(Resources.ProductsResource(products))
}

// Show single product details by id
func SingleProduct(c *gin.Context) {
	r := Application.NewRequest(c)

	var product Models.Product
	productId := c.Params.ByName("id")
	if err := r.DB.Where("id = ?", productId).First(&product).Error; err != nil {
		r.ResourceNotFound("product")
		return
	}
	r.Success(Resources.ProductResource(product))
}
