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
		log.Println("[CreateUser][1/3] Yugabyte gagal")
		return err
	}

	if err := m.userDBRepo.InsertUserRedis(user); err != nil {
		log.Println("[CreateUser][2/3] Redis gagal")
		return err
	}

	users, err := m.userDBRepo.GetUserListYugabyte()
	if err != nil {
		log.Println("[CreateUser] Gagal mengambil list user")
	}

	if err := m.userDBRepo.InsertUserRedisBulk(users); err != nil {
		log.Println("[CreateUser][3/3] Memperbarui data di redis gagal")
		return err
	}

	return nil
}
