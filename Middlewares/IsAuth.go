package Middlewares

import (
	"GoGin/Utils/Token"

	"github.com/gin-gonic/gin"
)

func IsAuth(c *gin.Context) bool {
	err := Token.TokenValid(c)
	if err != nil {
		return false
	}
	return true
}
