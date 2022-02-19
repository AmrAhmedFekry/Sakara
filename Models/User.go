package Models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string  `json:"username" gorm:"type:varchar(50) not null;unique_index"`
	Password string  `json:"password" gorm:"size:100;not null;"`
	Deposit  float64 `json:"deposit"`
	Role     string  `json:"role" gorm:"type:varchar(50)"`
}

// Genrate hashed password from user password input
func (user *User) GenerateHashFromPlainPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

// Verify password input with hashed password
func (user *User) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
