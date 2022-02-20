package Middlewares

import (
	"GoGin/Utils/Token"

	"github.com/gin-gonic/gin"
)

// Only for validate Token
// TODO: Make Function return same request or response NotAuth Response to make it more flexible
func IsAuth(c *gin.Context) bool {
	err := Token.TokenValid(c)
	if err != nil {
		return false
	}
	return true
}
