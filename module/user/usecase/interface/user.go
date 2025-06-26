package interfaces

import (
	"context"
	"my-go-app/module/user/model"
)

// Interface for repository
type UserRepository interface {
	FindByID(ctx context.Context, id int64) (*model.User, error)
}
