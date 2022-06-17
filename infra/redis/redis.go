package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

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

func NewRedisClient() *RedisClient {
	redis := newPool()
	client := RedisClient{
		Redis: redis,
	}
	fmt.Println("berhasail menginisiasi redis")
	return &client
}
