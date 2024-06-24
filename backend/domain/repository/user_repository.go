package repository

import (
	"context"
	"myapp/domain/model"
)

type UserRepository interface {
	GetByUserNameAndPassword(ctx context.Context, userName, password string) (model.User, error)
}
