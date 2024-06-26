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
func (c CommentRepositoryImpl) Create(ctx context.Context, postID int, userID int, body string) (int, error) {
	result, err := c.db.ExecContext(ctx,
		`
                INSERT INTO comments
                (
					comments.post_id,
					comments.user_id,
					comments.body
                )
                VALUES
                (
					?,
					?,
					?
                )
        `, postID, userID, body)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// GetByID implements repository.CommentRepository.
func (c CommentRepositoryImpl) GetByID(ctx context.Context, id int) (model.Comment, error) {
	comment := dao.CommentTable{}
	if err := c.db.GetContext(ctx, &comment, `
		SELECT
			id,
			post_id,
			user_id,
			body,
			created_at,
			updated_at
		FROM
			comments
		WHERE
			id=?
	`, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Comment{}, nil
		}

		log.Println(err)

		return model.Comment{}, err
	}

	return dao.ConvertCommentTableToDomainComment(comment), nil
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
func (c CommentRepositoryImpl) Update(ctx context.Context, int int, userID int, commentID int, body string) error {
	_, err := c.db.ExecContext(ctx,
		`
									UPDATE comments SET body = ? WHERE id = ? 
									
					`, body, commentID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// delete
func (c CommentRepositoryImpl) Delete(ctx context.Context, commentID int) error {
	_, err := c.db.ExecContext(ctx,
		`
		DELETE FROM comments WHERE id = ? 
	`, commentID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
