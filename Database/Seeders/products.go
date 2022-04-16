package Seeders

import (
	"GoGin/Models"

	"gorm.io/gorm"
)

// Use go faker to seed data
func seedUsers(DB *gorm.DB) {
	DB.Create(&Models.Product{})
}
