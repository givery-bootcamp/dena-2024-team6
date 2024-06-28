//go:generate mockgen -source=post_repository.go -destination=post_repository_mock.go -package=repository
package repository

import (
	"context"
)

type LikeRepository interface {
	Create(ctx context.Context, postID int) (int, error)
	Update(ctx context.Context, postID int) error
	Close(ctx context.Context) error
	Get(ctx context.Context, postID int) (int, error)
}
