package user

import (
	"github.com/google/uuid"
	"log"
	"mini-project/entity"
)

func (m Manager) CreateUser(user entity.User) error {
	uuid, err := uuid.NewUUID()
	if err != nil {
		log.Println("Gagal membuat UUID")
		return err
	}
	user.UUID = uuid

	if err := m.userDBRepo.InsertUserYugabyte(user); err != nil {
		log.Println("[CreateUser][1/2] Yugabyte gagal")
		return err
	}

	return nil
}
