package repository

import (
	"context"
	"my-go-app/module/user/model"
)

func (r *postgresUserRepo) FindByID(ctx context.Context, id int64) (*model.User, error) {
	// dummy data
	return &model.User{
		ID:    id,
		Name:  "John Doe",
		Email: "john@example.com",
	}, nil
}
