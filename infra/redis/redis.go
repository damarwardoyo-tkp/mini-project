package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

//type RedisClient struct {
//	Redis *redis.Client
//}

type RedisClient struct {
	Redis *redis.Pool
}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

//func newConn() (*redis.Client, error) {
//	client := redis.NewClient(&redis.Options{
//		Addr:     "localhost:6379",
//		Password: "",
//		DB:       0,
//	})
//
//	if err := client.Ping(context.Background()).Err(); err != nil {
//		log.Printf("[Redis]Gagal menginisiasi koneksi ke redis, err: %v", err)
//		return nil, err
//	}
//	return client, nil
//}

func NewRedisClient() *RedisClient {
	redis := newPool()
	//if err != nil {
	//	log.Fatalln("[RedisClient]Gagal membuat Redis Client")
	//}
	client := RedisClient{
		Redis: redis,
	}
	fmt.Println("berhasail menginisiasi redis")
	return &client
}
