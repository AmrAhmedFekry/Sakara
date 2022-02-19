package Sellers

import (
	"GoGin/Application"
	"GoGin/Models"
	Resources "GoGin/Resources/Sellers"
	"GoGin/Validations/Sellers"

	"github.com/gin-gonic/gin"
)

// Create new product and authenticated seller id
func CreateProduct(c *gin.Context) {
	r := Application.NewRequestWithAuth(c)

	if r.IsSeller == false {
		r.NotAuth()
		return
	}

	// Binding Request Body to user Model
	var product Models.Product
	r.Context.ShouldBindJSON(&product)
	// Validate user Model
	if r.ValidateRequest(Sellers.InsertValidation(product)).FailsValidation() {
		return
	}
	product.SellerID = r.User.ID
	r.DB.Create(&product)
	r.Created(Resources.ProductResource(product))
}

// Update Product by id and authenticated seller id
func UpdateProduct(c *gin.Context) {
	r := Application.NewRequestWithAuth(c)
	if r.IsSeller == false {
		r.NotAuth()
		return
	}

	// Binding Request Body to user Model
	var product Models.Product
	productId := c.Params.ByName("id")
	if err := r.DB.Where("id = ?", productId).First(&product).Error; err != nil {
		r.ResourceNotFound("product")
		return
	}

	if product.SellerID != r.User.ID {
		r.NotAuth()
		return
	}

	r.Context.ShouldBindJSON(&product)
	// Validate user Model
	if r.ValidateRequest(Sellers.InsertValidation(product)).FailsValidation() {
		return
	}
	r.DB.Save(&product)
	r.Success(Resources.ProductResource(product))
}

// Delete Product by id and by authenticated seller id
func DeleteProduct(c *gin.Context) {
	r := Application.NewRequestWithAuth(c)
	if r.IsSeller == false {
		r.NotAuth()
		return
	}

	var product Models.Product
	productId := c.Params.ByName("id")

	if err := r.DB.Where("id = ?", productId).First(&product).Error; err != nil {
		r.ResourceNotFound("product")
		return
	}
	if product.SellerID != r.User.ID {
		r.NotAuth()
		return
	}
	r.DB.Delete(&product)
	r.ProductDeleted()
}

// Show Single Product by id and by authenticated seller id
func ShowProduct(c *gin.Context) {
	r := Application.NewRequestWithAuth(c)

	if r.IsSeller == false {
		r.NotAuth()
		return
	}

	var product Models.Product
	productId := c.Params.ByName("id")
	if err := r.DB.Where("id = ?", productId).First(&product).Error; err != nil {
		r.ResourceNotFound("product")
		return
	}
	r.Success(Resources.ProductResource(product))
}

// Get all products from database of authenticated seller
func ListProducts(c *gin.Context) {
	r := Application.NewRequestWithAuth(c)

	if r.IsSeller == false {
		r.NotAuth()
		return
	}
	var products []Models.Product
	r.DB.Find(&products)
	r.Success(Resources.ProductsResource(products))
}
