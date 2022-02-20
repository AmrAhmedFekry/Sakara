package Visitors

import (
	"GoGin/Models"
	"GoGin/Utils/Token"
)

/*
	The Idea behind this snippet is to use this function to return only required fields
	of every model in case of single user, list of users and depend on the API requested Actor
*/

// Map of  user data in case of single user and if there are common fields
// Between singel and list of users then use this function to set them
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

// Map of users data
func UsersResource(users []Models.User) []map[string]interface{} {
	mappedUsers := make([]map[string]interface{}, 0)
	for _, user := range users {
		mappedUsers = append(mappedUsers, UserResource(user, false))
	}
	return mappedUsers
}
