package user

import (
	"mini-project/entity"
	"mini-project/infra/nsq"
	"mini-project/module/user/repo"
)

type UserManager interface {
	CreateUser(request entity.UserRequest) (string, error)
	GetUserList() ([]byte, error)
	GetUser(string) ([]byte, error)
}

type Manager struct {
	userDBRepo  *repo.UserDBRepoImpl
	nsqProducer *nsq.NSQProducer
}

func NewUserManager(userRepo *repo.UserDBRepoImpl, nsqProducer *nsq.NSQProducer) UserManager {
	manager := Manager{
		userDBRepo:  userRepo,
		nsqProducer: nsqProducer,
	}
	return &manager
}
