package user

import (
	"encoding/json"
	"log"
)

func (m Manager) GetUserList() (string, error) {
	usersRedis, err := m.userDBRepo.GetUserListRedis()
	if usersRedis == "" {
		users, err := m.userDBRepo.GetUserListYugabyte()
		if err != nil {
			log.Println("[GetUserList] Gagal mendapatkan list user dari yugabyte dan redis")
			return "", err
		}
		m.userDBRepo.InsertUserRedisBulk(users)
		resp, err := json.Marshal(users)
		return string(resp), err
	}
	return usersRedis, err
}
