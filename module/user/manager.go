package user

import (
	"context"
	"mini-project/entity"
	"mini-project/infra/db"
	"mini-project/infra/redis"
	"mini-project/module/user/repo"
)

type UserManager interface {
	CreateUser(context context.Context, user entity.User) error
	GetUserList()
	GetUser()
}

type Manager struct {
	userDBRepo *repo.UserDBRepoImpl
}

func NewUserManager(redis *redis.RedisClient, yugabyte *db.YugabyteClient) UserManager {
	userRepo := repo.NewUserDBRepo(redis, yugabyte)
	a := Manager{
		userDBRepo: userRepo,
	}
	return &a
}
