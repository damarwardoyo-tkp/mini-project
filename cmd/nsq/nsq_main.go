package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"mini-project/infra/db"
	"mini-project/infra/es"
	"mini-project/infra/redis"
	"mini-project/module/user/repo"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	redisClient := redis.NewRedisClient()
	yugaByteClient := db.NewYugabyteClient()
	esClient := es.NewElasticsearchClient()

	userRepo := repo.NewUserDBRepo(redisClient, yugaByteClient, esClient)

	redisMessageHandler := redisMessageHandler{
		userDBRepo: userRepo,
	}
	yugabyteMessageHandler := yugabyteMessageHandler{
		userDBRepo: userRepo,
	}
	esMessageHandler := esMessageHandler{
		userDBRepo: userRepo,
	}

	// Instantiate a consumer that will subscribe to the provided channel.
	config := nsq.NewConfig()
	consumerES, err := nsq.NewConsumer("test", "elasticsearch", config)
	if err != nil {
		log.Fatal(err)
	}
	consumerRedis, err := nsq.NewConsumer("test", "redis", config)
	if err != nil {
		log.Fatal(err)
	}
	consumerYuga, err := nsq.NewConsumer("test", "yugabyte", config)
	if err != nil {
		log.Fatal(err)
	}

	consumerES.AddHandler(&esMessageHandler)
	consumerYuga.AddHandler(&yugabyteMessageHandler)
	consumerRedis.AddHandler(&redisMessageHandler)

	consumerES.ConnectToNSQLookupd("localhost:4161")
	consumerRedis.ConnectToNSQLookupd("localhost:4161")
	consumerYuga.ConnectToNSQLookupd("localhost:4161")

	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	consumerES.Stop()
	consumerRedis.Stop()
	consumerYuga.Stop()

}
