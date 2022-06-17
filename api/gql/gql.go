package gql

import "mini-project/module/user"

type GqlHandler struct {
	manager user.UserManager
}

func NewGqlHandler(userManager user.UserManager) *GqlHandler {
	handler := GqlHandler{
		manager: userManager,
	}
	return &handler
}
