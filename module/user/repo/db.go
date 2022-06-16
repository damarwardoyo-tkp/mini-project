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
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func (repo *UserDBRepoImpl) InsertUserRedis(user entity.User) error {
	client := repo.redisClient.Redis.Get()
	defer client.Close()

	userJson, err := json.Marshal(&user)
	if err != nil {
		log.Printf("error ketika marshal data, err: %v", err)
		return err

	}
	_, err = client.Do("SETEX", user.Nama, 300, string(userJson))
	if err != nil {
		log.Printf("[Redis] Gagal insert data user %s ke redis, err: %s", user.Nama, err)
		return err
	}
	return nil
}

func (repo *UserDBRepoImpl) GetUserYugabyte(nama string) (entity.User, error) {
	var user entity.User
	if result := repo.yugabyteClient.DB.Where("nama = ?", nama).First(&user); result.Error != nil {
		log.Printf("[Yugabyte] Gagal mengambil data user %s dari yugabyte, err: %s", nama, result.Error)
		return user, result.Error
	}
	return user, nil
}

func (repo *UserDBRepoImpl) GetUserRedis(nama string) (string, error) {
	client := repo.redisClient.Redis.Get()
	defer client.Close()

	value, err := client.Do("GET", nama)
	if value == nil {
		log.Printf("[Redis] Gagal mengambil data user %s dari redis, err: %v", nama, err)
		return "", err
	}
	return string(value.([]uint8)), nil
}

func (repo *UserDBRepoImpl) GetUserListYugabyte() ([]entity.User, error) {
	var users []entity.User
	repo.yugabyteClient.DB.Find(&users)

	if result := repo.yugabyteClient.DB.Find(&users); result.Error != nil {
		log.Printf("[Yugabyte] Gagal mengambil data list user dari yugabyte, err: %s", result.Error)
		return nil, result.Error
	}
	return users, nil
}

func (repo *UserDBRepoImpl) GetUserListRedis() {
	client := repo.redisClient.Redis.Get()
	defer client.Close()

	_, err := client.Do("GET", 0)
	if err != nil {
		log.Printf("[Redis] Gagal mengambil data list user dari redis, err: %v", err)
		//return "", err
	}
}
