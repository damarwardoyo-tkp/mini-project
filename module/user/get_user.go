package user

import (
	"encoding/json"
	"log"
)

func (m Manager) GetUser(nama string) (string, error) {
	userRedis, _ := m.userDBRepo.GetUserRedis(nama)
	if userRedis == "" {
		user, err := m.userDBRepo.GetUserYugabyte(nama)
		if err != nil {
			log.Printf("[GetUser] Gagal mendapatkan user %v dari yugabyte dan redis", nama)
			return "", err
		}
		m.userDBRepo.InsertUserRedis(user)
		resp, err := json.Marshal(user)
		return string(resp), err
	}
	return userRedis, nil

}
