package user

import (
	"context"
	"github.com/google/uuid"
	"log"
	"mini-project/entity"
)

func (m Manager) CreateUser(ctx context.Context, user entity.User) error {
	uuid, err := uuid.NewUUID()
	if err != nil {
		log.Println("Gagal membuat UUID")
		return err
	}
	user.UUID = uuid

	if err := m.userDBRepo.InsertUserYugabyte(user); err != nil {
		log.Println("[CreateUser][1/2] Yugabyte gagal")
		return err
	}

	if err := m.userDBRepo.InsertUserRedis(ctx, user); err != nil {
		log.Println("[CreateUser][2/2] Redis gagal")
		return err
	}

	return nil
}
