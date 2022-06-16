package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"log"
)

type RedisClient struct {
	Redis *redis.Client
}

func newConn() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Printf("[Redis]Gagal menginisiasi koneksi ke redis, err: %v", err)
		return nil, err
	}
	return client, nil
}

func NewRedisClient() *RedisClient {
	redis, err := newConn()
	if err != nil {
		log.Fatalln("[RedisClient]Gagal membuat Redis Client")
	}
	client := RedisClient{
		Redis: redis,
	}
	fmt.Println("berhasail menginisiasi redis")
	return &client
}
