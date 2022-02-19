package Buyers

import (
	"GoGin/Application"
	"GoGin/Models"

	"github.com/gin-gonic/gin"
)

func Buy(c *gin.Context) {
	r := Application.NewRequestWithAuth(c)
	if r.IsBuyer == false {
		r.NotAuth()
		return
	}

	productId := c.Params.ByName("product_id")
	// requiredAmount := c.Params.ByName("product_amount")

	var product Models.Product
	if err := r.DB.Where("id = ?", productId).First(&product).Error; err != nil {
		r.ResourceNotFound("product")
		return
	}

	// if product.AmountAvailable < int(s
}
