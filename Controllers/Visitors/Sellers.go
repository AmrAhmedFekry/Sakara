package Visitors

import (
	"GoGin/Application"
	"GoGin/Models"
	Resources "GoGin/Resources/Visitors"
	"GoGin/Validations/Visitors"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Create New User
func RegisterAsSeller(c *gin.Context) {
	r := Application.NewRequest(c)
	// Binding Request Body to user Model
	var user Models.User
	r.Context.ShouldBindJSON(&user)
	// Validate user Model
	if r.ValidateRequest(Visitors.RegisterValidation(user)).FailsValidation() {
		return
	}
	r.DB.Where("user_name = ?", user.UserName).First(&user)
	if user.ID != 0 {
		r.ResourceAlreadyExists("user_name")
		return
	}
	user.Role = "seller"
	user.Password = user.GenerateHashFromPlainPassword(user.Password)
	r.DB.Create(&user)
	r.Created(Resources.UserResource(user, false))
}

// Login User
func LoginAsSeller(c *gin.Context) {
	r := Application.NewRequest(c)
	// Binding Request Body to user Model
	var user Models.User
	r.Context.ShouldBindJSON(&user)
	userPassword := user.Password
	if r.ValidateRequest(Visitors.LoginValidation(user)).FailsValidation() {
		return
	}

	r.DB.Where("user_name = ? AND role = ?", user.UserName, "seller").First(&user)

	if user.ID == 0 {
		r.ResourceNotFound("user")
		return
	}
	err := user.VerifyPassword(user.Password, userPassword)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		r.ResourceNotFound("user")
		return
	}
	r.Success(Resources.UserResource(user, true))
}
