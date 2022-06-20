package user

import (
	"mini-project/entity"
	"mini-project/infra/db"
	"mini-project/infra/redis"
	"mini-project/module/user/repo"
)

type UserManager interface {
	CreateUser(entity.User) (string, error)
	GetUserList() (string, error)
	GetUser(string) (string, error)
}

type Manager struct {
	userDBRepo *repo.UserDBRepoImpl
}

func NewUserManager(redis *redis.RedisClient, yugabyte *db.YugabyteClient) UserManager {
	userRepo := repo.NewUserDBRepo(redis, yugabyte)
	manager := Manager{
		userDBRepo: userRepo,
	}
	return &manager
}
