package repository

import (
	"context"
	"myapp/domain/model"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int) (model.User, error)
	GetByUserNameAndPassword(ctx context.Context, userName, password string) (model.User, error)
}
