package graph

import "mini-project/module/user"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	manager user.UserManager
}

func NewGQLHandler(manager user.UserManager) *Resolver {
	handler := Resolver{
		manager: manager,
	}
	return &handler
}
