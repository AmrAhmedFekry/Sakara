package Routes

import "GoGin/Application"

type RouterApp struct {
	*Application.Bootstrap
}

func (app RouterApp) Routing() {
	app.ExposedRoutes()
	app.productRoutes()
}
