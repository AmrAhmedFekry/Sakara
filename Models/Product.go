package Models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	AmountAvailable int32   `json:"amount_available"`
	Cost            float32 `json:"cost"`
	ProductName     string  `json:"product_name" gorm:"type:varchar(50)"`
	SellerID        uint    `json:"seller_id"`
}
