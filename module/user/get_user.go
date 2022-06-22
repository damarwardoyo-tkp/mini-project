package user

import (
	"encoding/json"
	"log"
	"mini-project/entity"
)

func (m Manager) GetUser(searchable string) ([]byte, error) {
	searchResult, err := m.userDBRepo.SearchUserES(searchable)
	if err != nil {
		log.Printf("[GetUser] Gagal mendapatkan user %v dari elasticsearch", searchable)
	}
	log.Println(len(searchResult))
	if len(searchResult) > 0 {
		userListEs, err := json.Marshal(searchResult)
		if err != nil {
			log.Printf("[GetUser] Gagal marshal data user, err:%v", err)
		} else {
			log.Println("[GetUser] dari es")
			return userListEs, err
		}
	}

	var usersRedis []interface{}
	var usersYuga []entity.User

	for _, item := range searchResult {
		if user, ok := item.(entity.User); ok {
			userRedis, _ := m.userDBRepo.GetUserRedisByUUID(user.UUID.String())
			if userRedis == nil {
				user, err := m.userDBRepo.GetUserYugabyteByUUID(user.UUID.String())
				if err != nil {
					log.Printf("[GetUser] Gagal mendapatkan user %v dari yugabyte dan redis", err)
				}
				usersYuga = append(usersYuga, user)
			} else {
				usersRedis = append(usersRedis, userRedis)
			}
		}
	}

	if len(usersRedis) > 0 {
		resp, err := json.Marshal(usersRedis)
		if err != nil {
			log.Printf("[GetUserList] Gagal marshal data list user, err:%v", err)
			return nil, err
		}
		log.Println("[get user] dari redis")
		return resp, err
	} else if len(usersYuga) > 0 {
		resp, err := json.Marshal(usersYuga)
		if err != nil {
			log.Printf("[GetUserList] Gagal marshal data list user, err:%v", err)
			return nil, err
		}
		log.Println("[get user] dari yugabyte")
		return resp, err
	}

	users, err := m.userDBRepo.GetUserYugabyteBySearchable(searchable)
	if err != nil {
		log.Println("[GetUser] Gagal mendapatkan list user dari yugabyte dan redis")
	}
	resp, err := json.Marshal(users)
	if err != nil {
		log.Printf("[GetUser] Gagal marshal data list user, err:%v", err)
	}
	log.Println("[get user] dari yugabyte searchable")
	return resp, err

	//userRedis, _ := m.userDBRepo.GetUserRedisByUUID(searchable)
	//if userRedis == nil {
	//	user, err := m.userDBRepo.GetUserYugabyteByUUID(searchable)
	//	if err != nil {
	//		log.Printf("[GetUser] Gagal mendapatkan user %v dari yugabyte dan redis", err)
	//		return nil, err
	//	}
	//	m.userDBRepo.InsertUserRedis(user.UUID.String(), user)
	//	resp, err := json.Marshal(user)
	//	if err != nil {
	//		log.Printf("[GetUser] Gagal marshal data user, err:%v", err)
	//		return nil, err
	//	}
	//	return resp, err
	//}
	//return userRedis.([]byte), nil
}
