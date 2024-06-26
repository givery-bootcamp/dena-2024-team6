//go:generate mockgen -source=post_repository.go -destination=post_repository_mock.go -package=repository
package repository

import (
	"context"
	"myapp/domain/model"
)

type PostRepository interface {
	Create(ctx context.Context, userID int, title, body string) (int, error)
	List(ctx context.Context) ([]model.Post, error)
	GetDetail(ctx context.Context, id int) (model.PostDetail, error)

	Update(ctx context.Context, id int, title, body string) error
	Delete(ctx context.Context, id int) error
}
