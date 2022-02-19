package Routes

import "GoGin/Controllers/Visitors"

func (app RouterApp) ExposedRoutes() {
	// Login and register for seller|buyer routes
	app.Gin.POST("/api/seller/register", Visitors.RegisterAsSeller)
	app.Gin.POST("/api/seller/login", Visitors.LoginAsSeller)
	app.Gin.POST("/api/buyer/register", Visitors.RegisterAsBuyer)
	app.Gin.POST("/api/buyer/login", Visitors.LoginAsBuyer)

	// Product routes
	app.Gin.GET("/api/buyer/products", Visitors.ListOfProducts)
	app.Gin.GET("/api/buyer/product/:id", Visitors.SingleProduct)
}
