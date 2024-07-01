package repository

import (
	"context"
	"myapp/domain/model"
)

type UserRepository interface {
	Create(ctx context.Context, userName, password string) (model.User, error)
	GetByID(ctx context.Context, id int) (model.User, error)
	GetByUserNameAndPassword(ctx context.Context, userName, password string) (model.User, error)
	Exists(ctx context.Context, userName string) (bool, error)
}
