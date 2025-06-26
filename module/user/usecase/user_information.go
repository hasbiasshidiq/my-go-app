package usecase

import (
	"context"
	"my-go-app/module/user/model"
	interfaces "my-go-app/module/user/usecase/interface"
)

type UserUsecase struct {
	userRepo interfaces.UserRepository
}

func NewUserUsecase(repo interfaces.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: repo}
}

func (u *UserUsecase) GetByID(ctx context.Context, id int64) (*model.User, error) {
	return u.userRepo.FindByID(ctx, id)
}
