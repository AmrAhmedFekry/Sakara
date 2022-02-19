package Application

import (
	"GoGin/Database"
	"GoGin/Database/Seeders"
	"database/sql"
	"fmt"

	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Bootstrap struct {
	Gin        *gin.Engine
	DB         *gorm.DB
	Connection *sql.DB
}

func (app *Bootstrap) Share() {}

func app() func() *Bootstrap {
	return func() *Bootstrap {
		var app Bootstrap
		app.Gin = gin.Default()
		connectToDatabase(&app)
		err := gotrans.InitLocales("./Public/locales")
		if err != nil {
			fmt.Println("This is error in loading locale files: ", err.Error())
		}

		return &app
	}
}

func NewApp() *Bootstrap {
	app := app()
	return app()
}

// Migrate Database Models
func (app *Bootstrap) Migrate() {
	Database.Migrate(app.DB)
}

// Seed Database Models with data
func (app *Bootstrap) Seed() {
	Seeders.Seeds(app.DB)
}
