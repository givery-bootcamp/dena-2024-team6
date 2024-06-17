//go:generate mockgen -source=post_repository.go -destination=post_repository_mock.go -package=repository
package repository

import (
	"context"
	"myapp/core/model"
)

type PostRepository interface {
	List(ctx context.Context) ([]model.Post, error)
	GetDetail(ctx context.Context, id int) (model.Post, error)
}
