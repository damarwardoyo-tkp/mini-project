package repo

import (
	"context"
	"encoding/json"
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

func (repo *UserDBRepoImpl) InsertUserRedis(ctx context.Context, user entity.User) error {
	uuid := user.UUID.String()
	userJson, err := json.Marshal(&user)
	if err != nil {
		log.Printf("error ketika marshal data, err: %v", err)
		return err
	}
	err = repo.redisClient.Redis.Set(ctx, uuid, string(userJson), 300).Err()
	if err != nil {
		log.Printf("[Redis] Gagal insert data user ke redis, err: %v", err)
		return err
	}
	return nil
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
