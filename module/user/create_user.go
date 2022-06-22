package user

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"mini-project/entity"
)

func (m Manager) CreateUser(userReq entity.UserRequest) (string, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		log.Println("Gagal membuat UUID")
		return "", err
	}
	user := entity.User{
		UUID:       uuid,
		Nama:       userReq.Nama,
		Umur:       userReq.Umur,
		Alamat:     userReq.Alamat,
		Searchable: fmt.Sprintf(uuid.String()+" "+userReq.Nama+" "+userReq.Alamat+" "+"%d", userReq.Umur),
	}
	userMessage := entity.UserMessage{
		Request: "create",
		User:    user,
	}
	payload, err := json.Marshal(userMessage)
	if err != nil {
		return "", err
	}
	m.nsqProducer.Producer.Publish("test", payload)

	//
	//if err := m.userDBRepo.InsertUserES(user); err != nil {
	//	log.Println("[CreateUser][1/4] Elasticsearch gagal")
	//	return "", err
	//}
	//
	//if err := m.userDBRepo.InsertUserYugabyte(user); err != nil {
	//	log.Println("[CreateUser][2/4] Yugabyte gagal")
	//	return "", err
	//}
	//
	//if err := m.userDBRepo.InsertUserRedis(user.UUID.String(), user); err != nil {
	//	log.Println("[CreateUser][3/4] Redis gagal")
	//	return "", err
	//}
	//
	//if err := m.userDBRepo.InsertUserRedis(user.Nama, user); err != nil {
	//	log.Println("[CreateUser][3/4] Redis gagal")
	//	return "", err
	//}
	//
	//users, err := m.userDBRepo.GetAllUserYugabyte()
	//if err != nil {
	//	log.Println("[CreateUser] Gagal mengambil list user")
	//}
	//
	//if err := m.userDBRepo.InsertUserRedisBulk(users); err != nil {
	//	log.Println("[CreateUser][4/4] Memperbarui data di redis gagal")
	//	return "", err
	//}

	return uuid.String(), err
}
