package user

import (
	"log"
	"mini-project/entity"
)

func (m Manager) GetUserList() ([]entity.User, error) {
	users, err := m.userDBRepo.GetUserListYugabyte()
	if err != nil {
		log.Println("[GetUserList] Gagal mengambil list user")
		return nil, err
	}
	return users, nil
}

func (m Manager) GetUserListRedis() (string, error) {
	//users, err := m.userDBRepo.GetUserListYugabyte()
	//if err != nil {
	//	log.Println("[GetUserList] Gagal mengambil list user")
	//	return "nil", err
	//}
	return "users", nil
}
