package rest

import "mini-project/module/user"

type RestHandler struct {
	manager user.UserManager
}

func NewRestHandler(userManager user.UserManager) *RestHandler {
	handler := RestHandler{
		manager: userManager,
	}
	return &handler
}
