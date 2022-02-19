package main

import (
	"GoGin/Application"
	"GoGin/Routes"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load(".env")

	app := Application.NewApp()
	// migrate models
	app.Migrate()
	// Seed data
	// app.Seed()

	Application.CloseConnection(app)
	routerApp := Routes.RouterApp{app}
	routerApp.Routing()
	app.Gin.Run(":8080")
}
