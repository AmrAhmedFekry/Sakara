package Routes

import (
	"GoGin/Controllers/Sellers"
)

// This routes only available for authenticated users with type seller
// Throw this APIs sellers can manage their products
func (app RouterApp) productRoutes() {
	app.Gin.POST("/api/create/product", Sellers.CreateProduct)
	app.Gin.PUT("/api/product/:id", Sellers.UpdateProduct)
	app.Gin.GET("/api/product/:id", Sellers.ShowProduct)
	app.Gin.DELETE("/api/product/:id", Sellers.DeleteProduct)
	app.Gin.GET("/api/products", Sellers.ListProducts)
}

// func (app RouterApp) buyerRoutes() {
// 	app.Gin.POST("/api/buy", Buyers.Buy)
// }
