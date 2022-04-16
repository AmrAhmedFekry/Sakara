package Routes

import "GoGin/Controllers/Visitors"

func (app RouterApp) ExposedRoutes() {
	// Product routes
	app.Gin.GET("/api/products", Visitors.ListOfProducts)
	app.Gin.GET("/api/product/:id", Visitors.SingleProduct)
}
