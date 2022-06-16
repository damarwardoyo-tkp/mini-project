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
			log.Printf("[GetUser] Gagal menemukan user %v dari yugabyte untuk memperbaharui cache. err:%v", nama, err)
			return "", err
		}
		m.userDBRepo.InsertUserRedis(user)
		log.Println("[GetUser] memperbarui data di redis")
		resp, _ := json.Marshal(user)
		return string(resp), nil
	}
	return userRedis, nil

}
