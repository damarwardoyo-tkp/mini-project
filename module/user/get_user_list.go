package user

import (
	"encoding/json"
	"fmt"
	"log"
)

func (m Manager) GetUserList() (string, error) {
	usersRedis, err := m.userDBRepo.GetUserListRedis()
	if usersRedis == "" {
		users, err := m.userDBRepo.GetUserListYugabyte()
		if err != nil {
			log.Println("[GetUserList] Gagal mengambil list user dari yugabyte")
			return "", err
		}
		m.userDBRepo.InsertUserRedisBulk(users)
		resp, err := json.Marshal(users)
		fmt.Println("data user list dari yugabyte")
		return string(resp), err
	}
	fmt.Println("data user list dari redis")
	return usersRedis, err
}
