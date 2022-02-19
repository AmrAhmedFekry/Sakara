package Database

import (
	"GoGin/Models"

	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(&Models.User{}, &Models.Product{})
}
