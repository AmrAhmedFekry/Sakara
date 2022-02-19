package Seeders

import "gorm.io/gorm"

func Seeds(DB *gorm.DB) {
	seedUsers(DB)
}
