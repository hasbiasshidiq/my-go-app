package interfaces

import (
	"context"
	"my-go-app/module/user/model"
)

type UserUsecase interface {
	GetByID(ctx context.Context, id int64) (*model.User, error)
}
