package repository

import (
	"context"
	"myapp/domain/model"
)

type CommentRepository interface {
	List(ctx context.Context, postID int) ([]model.Comment, error)
	GetByID(ctx context.Context, id int) (model.Comment, error)

	Create(ctx context.Context, postID int, userID int, body string) error
	Update(ctx context.Context, postID int, userID int, commentID int, body string) error
	Delete(ctx context.Context, postID int, commentID int) error
}
