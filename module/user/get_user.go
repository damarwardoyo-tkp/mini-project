package user

import "fmt"

func (m Manager) GetUser() {
	fmt.Println("get an user")
	m.userDBRepo.GetUserYugabyte()
}
