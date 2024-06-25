package store

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"myapp/domain/model"
	"myapp/domain/repository"
	"myapp/infrastructure/store/dao"

	"github.com/jmoiron/sqlx"
	"github.com/samber/do"
)

type CommentRepositoryImpl struct {
	db *sqlx.DB
}

func NewCommentRepositoryImpl(i *do.Injector) (repository.CommentRepository, error) {
	db := do.MustInvoke[*sqlx.DB](i)
	return CommentRepositoryImpl{
		db: db,
	}, nil
}

// CreateComment implements repository.CommentRepository.
func (c CommentRepositoryImpl) CreateComment(ctx context.Context, int int, userID int, body string) (model.Comment, error) {
	panic("unimplemented")
}

// DeleteComment implements repository.CommentRepository.
func (c CommentRepositoryImpl) DeleteComment(ctx context.Context, int int, commentID int) error {
	panic("unimplemented")
}

// GetByID implements repository.CommentRepository.
func (c CommentRepositoryImpl) GetByID(ctx context.Context, id int) (model.Comment, error) {
	panic("unimplemented")
}

// List implements repository.CommentRepository.
func (c CommentRepositoryImpl) List(ctx context.Context, postID int) ([]model.Comment, error) {
	comments := []dao.CommentTable{}
	if err := c.db.SelectContext(ctx, &comments, `
		SELECT
			comments.id,
			comments.post_id,
			comments.user_id,
			comments.body,
			comments.created_at,
			comments.updated_at,
			users.name as "user_name"
		FROM
			comments
		INNER JOIN
			users
		ON
			comments.user_id=users.id
		WHERE
			comments.post_id=?
	`, postID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		log.Println(err)

		return nil, err
	}

	// テーブルのモデルからドメインモデルに変換
	domainComments := make([]model.Comment, len(comments))
	for i, com := range comments {
		domainComments[i] = dao.ConvertCommentTableToDomainComment(com)
	}
	return domainComments, nil
}

// UpdateComment implements repository.CommentRepository.
func (c CommentRepositoryImpl) UpdateComment(ctx context.Context, int int, userID int, commentID int, body string) (model.Comment, error) {
	panic("unimplemented")
}
