package user

import (
	"encoding/json"
	"log"
)

func (m Manager) GetUserList() ([]byte, error) {
	userListEsRAW, err := m.userDBRepo.GetAllUserES()
	if err != nil {
		log.Println("[GetUserList] Gagal mendapatkan list user dari elasticsearch")
	}

	if len(userListEsRAW) > 0 {
		userListEs, err := json.Marshal(userListEsRAW)
		if err != nil {
			log.Printf("[GetUserList] Gagal marshal data list user, err:%v", err)
		} else {
			log.Println("[GetUserList] dari es")
			return userListEs, err
		}
	}

	userListRedis, err := m.userDBRepo.GetAllUserRedis()
	if userListRedis == nil {
		users, err := m.userDBRepo.GetAllUserYugabyte()
		if err != nil {
			log.Println("[GetUserList] Gagal mendapatkan list user dari yugabyte dan redis")
			return nil, err
		}
		m.userDBRepo.InsertUserRedisBulk(users)
		resp, err := json.Marshal(users)
		if err != nil {
			log.Printf("[GetUserList] Gagal marshal data list user, err:%v", err)
			return nil, err
		}
		return resp, err
	}
	return userListRedis.([]byte), err
}
