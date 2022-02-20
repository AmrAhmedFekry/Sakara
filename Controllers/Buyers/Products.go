package Buyers

import (
	"GoGin/Application"
	"GoGin/Models"
	"encoding/json"
	"io/ioutil"

	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
)

// Buy Product with specified product id and amount
// Check if product exits and with enough amount,
// Check if total price required to buy these amount meet the buyer balance
func Buy(c *gin.Context) {
	r := Application.NewRequestWithAuth(c)
	if r.IsBuyer == false {
		r.NotAuth()
		return
	}
	type BuyRequest struct {
		ProductID int32 `json:"product_id"`
		Amount    int32 `json:"amount"`
	}
	var jsonData BuyRequest
	data, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal([]byte(data), &jsonData)

	// Check if there is product available and if the amount is available
	var product Models.Product
	if err := r.DB.Where("id = ? ", jsonData.ProductID).First(&product).Error; err != nil {
		r.ResourceNotFound("product")
		return
	}
	// Check if the amount is available
	if product.AmountAvailable < jsonData.Amount {
		r.CustomResponse(gotrans.T("product_not_available"), 400)
		return
	}
	var Buyer Models.User
	r.DB.Where("id = ?", r.User.ID).First(&Buyer)
	productAmountCost := float32(jsonData.Amount) * product.Cost
	// Check if the total price is available in buyer deposit
	if Buyer.Deposit < float64(productAmountCost) {
		r.CustomResponse(gotrans.T("insufficient_balance"), 400)
		return
	}

	product.AmountAvailable -= jsonData.Amount
	r.DB.Save(&product)

	Buyer.Deposit -= float64(productAmountCost)
	r.DB.Save(&Buyer)

	r.Success(gotrans.T("transaction_success"))
	return
}
