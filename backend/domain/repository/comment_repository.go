package repository

import (
	"context"
	"myapp/domain/model"
)

type CommentRepository interface {
	List(ctx context.Context, postID int) ([]model.Comment, error)
	GetByID(ctx context.Context, id int) (model.Comment, error)

	CreateComment(ctx context.Context, int, userID int, body string) (model.Comment, error)
	UpdateComment(ctx context.Context, int, userID int, commentID int, body string) (model.Comment, error)
	DeleteComment(ctx context.Context, int, commentID int) error
}
