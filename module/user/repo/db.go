package repo

import (
	"fmt"
	"log"
	"mini-project/entity"
	"mini-project/infra/db"
	"mini-project/infra/redis"
)

type UserDBRepoImpl struct {
	yugabyteClient *db.YugabyteClient
	redisClient    *redis.RedisClient
}

func NewUserDBRepo(redis *redis.RedisClient, yugabyte *db.YugabyteClient) *UserDBRepoImpl {
	userDBRepo := UserDBRepoImpl{
		yugabyteClient: yugabyte,
		redisClient:    redis,
	}
	return &userDBRepo
}

func (repo *UserDBRepoImpl) InsertUserYugabyte(user entity.User) error {
	if result := repo.yugabyteClient.DB.Create(&user); result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func (repo *UserDBRepoImpl) InsertUserRedis(user entity.User) {
	fmt.Printf("nama ku %v", user.Nama)
}

func (repo *UserDBRepoImpl) GetUserYugabyte() {
	fmt.Println("user")
}

func (repo *UserDBRepoImpl) GetUserRedis() {
	fmt.Println("user")
}

func (repo *UserDBRepoImpl) GetUserListYugabyte() {
	fmt.Println("list user")
}

func (repo *UserDBRepoImpl) GetUserListRedis() {
	fmt.Println("list user")
}
