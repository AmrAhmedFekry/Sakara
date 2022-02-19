package Visitors

import (
	"GoGin/Models"
	"GoGin/Utils/Token"
)

func UserResource(user Models.User, isRequiredToken bool) map[string]interface{} {
	userResource := make(map[string]interface{})
	userResource["username"] = user.UserName
	userResource["id"] = user.ID
	userResource["deposit"] = user.Deposit
	userResource["role"] = user.Role
	if isRequiredToken {
		token, _ := Token.GenerateToken(user.ID)
		userResource["token"] = token
	}
	return userResource
}

func UsersResource(users []Models.User) []map[string]interface{} {
	mappedUsers := make([]map[string]interface{}, 0)
	for _, user := range users {
		mappedUsers = append(mappedUsers, UserResource(user, false))
	}
	return mappedUsers
}
