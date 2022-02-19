package Visitors

import (
	"GoGin/Application"
	"GoGin/Models"
	Resources "GoGin/Resources/Visitors"
	"GoGin/Validations/Visitors"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Create New User as buyer
func RegisterAsBuyer(c *gin.Context) {
	r := Application.NewRequest(c)
	// Binding Request Body to user Model
	var user Models.User
	r.Context.ShouldBindJSON(&user)
	// Validate user Model
	if r.ValidateRequest(Visitors.RegisterValidation(user)).FailsValidation() {
		return
	}
	r.DB.Where("user_name = ? ", user.UserName).First(&user)
	if user.ID != 0 {
		r.ResourceAlreadyExists("user_name")
		return
	}
	user.Password = user.GenerateHashFromPlainPassword(user.Password)
	user.Role = "buyer"
	r.DB.Create(&user)
	r.Created(Resources.UserResource(user, false))
}

// Login User as buyer
func LoginAsBuyer(c *gin.Context) {
	r := Application.NewRequest(c)
	// Binding Request Body to user Model
	var user Models.User
	r.Context.ShouldBindJSON(&user)
	userPassword := user.Password
	if r.ValidateRequest(Visitors.LoginValidation(user)).FailsValidation() {
		return
	}

	r.DB.Where("user_name = ? AND role = ?", user.UserName, "buyer").First(&user)

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
