package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"log"
	"mini-project/entity"
	"mini-project/graph/generated"
	"mini-project/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	var user entity.User
	user.Nama = input.Nama
	user.Umur = input.Umur
	user.Alamat = input.Alamat
	user.Searchable = input.Searchable

	uuid, err := r.manager.CreateUser(user)
	if err != nil {
		log.Printf("[GQL]Gagal membuat user baru, err:%v", err)
	}
	var resp model.User
	resp.UUID = uuid
	resp.Nama = input.Nama
	resp.Umur = input.Umur
	resp.Alamat = input.Alamat
	resp.Searchable = input.Searchable

	return &resp, err
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var resp []*model.User
	users, err := r.manager.GetUserList()
	if err != nil || users == "[]" {
		log.Println(err)
	}
	if err := json.Unmarshal([]byte(users), &resp); err != nil {
		log.Printf("[GQL]Gagal mendapat data list user, err:%v", err)
	}
	return resp, nil
}

func (r *queryResolver) User(ctx context.Context, searchable string) (*model.User, error) {
	var resp *model.User
	user, err := r.manager.GetUser(searchable)
	if err != nil || user == "[]" {
		log.Println(err)
	}
	if err := json.Unmarshal([]byte(user), &resp); err != nil {
		log.Printf("[GQL]Gagal mendapat data user, err:%v", err)
	}
	return resp, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
