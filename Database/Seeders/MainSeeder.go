package Seeders

import "gorm.io/gorm"

// Initialize all seeders functions
func Seeds(DB *gorm.DB) {
	seedUsers(DB)
}
