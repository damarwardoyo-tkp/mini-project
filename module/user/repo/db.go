package repo

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"log"
	"mini-project/entity"
	"mini-project/infra/db"
	"mini-project/infra/es"
	"mini-project/infra/redis"
	"reflect"
)

type UserDBRepoImpl struct {
	yugabyteClient *db.YugabyteClient
	redisClient    *redis.RedisClient
	esClient       *es.ElasticsearchClient
}

func NewUserDBRepo(redis *redis.RedisClient, yugabyte *db.YugabyteClient, es *es.ElasticsearchClient) *UserDBRepoImpl {
	userDBRepo := UserDBRepoImpl{
		yugabyteClient: yugabyte,
		redisClient:    redis,
		esClient:       es,
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

func (repo *UserDBRepoImpl) InsertUserRedis(key string, user entity.User) error {
	client := repo.redisClient.Redis.Get()
	defer client.Close()

	userJson, err := json.Marshal(&user)
	if err != nil {
		log.Printf("[InsertUserRedis] Error ketika marshal data %s, err: %v", user.Nama, err)
		return err

	}
	_, err = client.Do("SETEX", key, 300, string(userJson))
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

func (repo *UserDBRepoImpl) GetUserYugabyteByUUID(uuid string) (entity.User, error) {
	var user entity.User
	if result := repo.yugabyteClient.DB.First(&user, "uuid = ?", uuid); result.Error != nil {
		log.Printf("[Yugabyte] Gagal mendapatkan data user %s dari yugabyte, err: %s", uuid, result.Error)
		return user, result.Error
	}
	return user, nil
}

func (repo *UserDBRepoImpl) GetUserYugabyteBySearchable(searchable string) ([]entity.User, error) {
	param := "%" + searchable + "%"
	var users []entity.User
	if result := repo.yugabyteClient.DB.Where("searchable like ?", param).Find(&users); result.Error != nil {
		log.Printf("[Yugabyte] Gagal mendapatkan data user %s dari yugabyte, err: %s", searchable, result.Error)
		return users, result.Error
	}
	return users, nil
}

func (repo *UserDBRepoImpl) GetUserRedisByUUID(uuid string) (interface{}, error) {
	client := repo.redisClient.Redis.Get()
	defer client.Close()

	value, err := client.Do("GET", uuid)
	if value == nil {
		log.Printf("[Redis] Gagal mendapatkan data user %s dari redis, err: %v", uuid, err)
		return nil, err
	}
	//value.([]byte)
	return value, nil
}

func (repo *UserDBRepoImpl) GetAllUserYugabyte() ([]entity.User, error) {
	var users []entity.User
	if result := repo.yugabyteClient.DB.Find(&users); result.Error != nil {
		log.Printf("[Yugabyte] Gagal mendapatkan data list user dari yugabyte, err: %s", result.Error)
		return nil, result.Error
	}
	return users, nil
}

func (repo *UserDBRepoImpl) GetAllUserRedis() (interface{}, error) {
	client := repo.redisClient.Redis.Get()
	defer client.Close()

	value, err := client.Do("GET", 0)
	if value == nil {
		log.Printf("[Redis] Gagal mendapatkan data list user dari redis, err: %v", err)
		return nil, err
	}
	//value.([]byte)
	return value, err
}

func (repo *UserDBRepoImpl) SearchUserES(searchable string) ([]interface{}, error) {
	// Search with a term query
	termQuery := elastic.NewMatchQuery("searchable", searchable).Analyzer("searchable_indexing").ZeroTermsQuery("all")
	searchResult, err := repo.esClient.ES.Search().
		Index("test3").          // search in index "tweets"
		Query(termQuery).        // specify the query
		Pretty(true).            // pretty print request and response JSON
		Do(context.Background()) // execute
	if err != nil {
		log.Printf("[Elasticsearch] Error ketika mencari document %v, err:%v,", searchable, err)
	}

	res := searchResult.Each(reflect.TypeOf(entity.User{}))
	//for _, item := range searchResult.Each(reflect.TypeOf(entity.User{})) {
	//	if t, ok := item.(entity.User); ok {
	//		fmt.Printf("Tweet by %s: %s\n", t, t.Searchable)
	//	}
	//}
	return res, err
}

func (repo *UserDBRepoImpl) GetAllUserES() ([]interface{}, error) {
	termQuery := elastic.MatchAllQuery{}
	searchResult, err := repo.esClient.ES.Search().
		Index("test3").          // search in index "tweets"
		Query(&termQuery).       // specify the query
		Pretty(true).            // pretty print request and response JSON
		Do(context.Background()) // execute
	if err != nil {
		log.Printf("[Elasticsearch] Gagal mendapat data. err:%v", err)
		return nil, err
	}

	//var users []entity.User
	userListRAW := searchResult.Each(reflect.TypeOf(entity.User{}))
	//userList, err := json.Marshal(userListRAW)
	//if err != nil {
	//	log.Printf("[GetUserList] Gagal marshal data list user, err:%v", err)
	//	return nil, err
	//}

	return userListRAW, err
}

func (repo *UserDBRepoImpl) InsertUserES(user entity.User) error {
	_, err := repo.esClient.ES.Index().
		Index("test3").
		Id(user.UUID.String()).
		BodyJson(user).
		Do(context.Background())
	if err != nil {
		log.Printf("[Elasticsearch] Gagal memasukan data %v ke index, err:%v", user.Nama, err)
	}
	return nil
}
