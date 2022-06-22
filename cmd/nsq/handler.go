package main

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"log"
	"mini-project/entity"
	"mini-project/module/user/repo"
)

type redisMessageHandler struct {
	userDBRepo *repo.UserDBRepoImpl
}

func (r redisMessageHandler) HandleMessage(message *nsq.Message) error {
	var request entity.UserMessage
	if err := json.Unmarshal(message.Body, &request); err != nil {
		log.Printf("[Redis HandleMessage] Gagal unmarshal request, err:%v", err)
		return err
	}

	if request.Request == "create" {
		if err := r.userDBRepo.InsertUserRedis(request.User.UUID.String(), request.User); err != nil {
			log.Println("[CreateUser][3/4] Redis gagal")
			return err
		}

		if err := r.userDBRepo.InsertUserRedis(request.User.Nama, request.User); err != nil {
			log.Println("[CreateUser][3/4] Redis gagal")
			return err
		}
	}
	return nil
}

type esMessageHandler struct {
	userDBRepo *repo.UserDBRepoImpl
}

func (e esMessageHandler) HandleMessage(message *nsq.Message) error {
	var request entity.UserMessage
	if err := json.Unmarshal(message.Body, &request); err != nil {
		log.Printf("[Redis HandleMessage] Gagal unmarshal request, err:%v", err)
		return err
	}

	if request.Request == "create" {
		if err := e.userDBRepo.InsertUserES(request.User); err != nil {
			log.Println("[CreateUser][1/4] Elasticsearch gagal")
			return err
		}
	}
	return nil
}

type yugabyteMessageHandler struct {
	userDBRepo *repo.UserDBRepoImpl
}

func (y yugabyteMessageHandler) HandleMessage(message *nsq.Message) error {
	var request entity.UserMessage
	if err := json.Unmarshal(message.Body, &request); err != nil {
		log.Printf("[Yugabyte HandleMessage] Gagal unmarshal request, err:%v", err)
		return err
	}
	if request.Request == "create" {
		if err := y.userDBRepo.InsertUserYugabyte(request.User); err != nil {
			log.Println("[CreateUser][2/4] Yugabyte gagal")
			return err
		}
	}
	return nil
}
