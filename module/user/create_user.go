package user

import (
	"fmt"
	"log"
	"mini-project/entity"
)

func (m Manager) CreateUser(user entity.User) error {
	fmt.Println("create user")
	if err := m.userDBRepo.InsertUserYugabyte(user); err != nil {
		log.Println("[CreateUser][1/2] Yugabyte gagal")
		return err
	}

	return nil
}
