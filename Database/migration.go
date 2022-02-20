package Database

import (
	"GoGin/Models"

	"gorm.io/gorm"
)

// Set your auto migrating model here
func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(&Models.User{}, &Models.Product{})
}
