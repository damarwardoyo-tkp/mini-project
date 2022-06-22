package user

import (
	"mini-project/entity"
	"mini-project/module/user/repo"
)

type UserManager interface {
	CreateUser(request entity.UserRequest) (string, error)
	GetUserList() ([]byte, error)
	GetUser(string) ([]byte, error)
}

type Manager struct {
	userDBRepo *repo.UserDBRepoImpl
}

func NewUserManager(userRepo *repo.UserDBRepoImpl) UserManager {
	manager := Manager{
		userDBRepo: userRepo,
	}
	return &manager
}
