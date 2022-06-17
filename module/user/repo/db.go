package repo

import (
	"encoding/json"
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
		log.Printf("[Redis] Gagal insert data user %s ke redis, err: %s", user.Nama, result.Error)
		return result.Error
	}
	return nil
}

func (repo *UserDBRepoImpl) InsertUserRedis(user entity.User) error {
	client := repo.redisClient.Redis.Get()
	defer client.Close()

	userJson, err := json.Marshal(&user)
	if err != nil {
		log.Printf("[InsertUserRedis] Error ketika marshal data %s, err: %v", user.Nama, err)
		return err

	}
	_, err = client.Do("SETEX", user.Nama, 300, string(userJson))
	if err != nil {
		log.Printf("[Redis] Gagal insert data user %s ke redis, err: %s", user.Nama, err)
		return err
	}
	return err
}

func (repo *UserDBRepoImpl) InsertUserRedisBulk(users []entity.User) error {
	client := repo.redisClient.Redis.Get()
	defer client.Close()

	usersJson, err := json.Marshal(&users)
	if err != nil {
		log.Printf("[InsertUserRedisBulk] Error ketika marshal data list user, err: %v", err)
		return err

	}
	_, err = client.Do("SETEX", 0, 300, string(usersJson))
	if err != nil {
		log.Printf("[Redis] Gagal insert data list user ke redis, err: %v", err)
		return err
	}
	return err
}

func (repo *UserDBRepoImpl) GetUserYugabyte(nama string) (entity.User, error) {
	var user entity.User
	if result := repo.yugabyteClient.DB.Where("nama = ?", nama).First(&user); result.Error != nil {
		log.Printf("[Yugabyte] Gagal mendapatkan data user %s dari yugabyte, err: %s", nama, result.Error)
		return user, result.Error
	}
	return user, nil
}

func (repo *UserDBRepoImpl) GetUserRedis(nama string) (string, error) {
	client := repo.redisClient.Redis.Get()
	defer client.Close()

	value, err := client.Do("GET", nama)
	if value == nil {
		log.Printf("[Redis] Gagal mendapatkan data user %s dari redis, err: %v", nama, err)
		return "", err
	}
	return string(value.([]uint8)), nil
}

func (repo *UserDBRepoImpl) GetUserListYugabyte() ([]entity.User, error) {
	var users []entity.User
	if result := repo.yugabyteClient.DB.Find(&users); result.Error != nil {
		log.Printf("[Yugabyte] Gagal mendapatkan data list user dari yugabyte, err: %s", result.Error)
		return nil, result.Error
	}
	return users, nil
}

func (repo *UserDBRepoImpl) GetUserListRedis() (string, error) {
	client := repo.redisClient.Redis.Get()
	defer client.Close()

	value, err := client.Do("GET", 0)
	if value == nil {
		log.Printf("[Redis] Gagal mendapatkan data list user dari redis, err: %v", err)
		return "", err
	}
	return string(value.([]uint8)), err
}
