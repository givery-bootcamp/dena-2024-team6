package store

import (
	"context"
	"database/sql"
	"errors"
	"myapp/domain/model"
	"myapp/domain/repository"
	"myapp/infrastructure/store/dao"

	"github.com/jmoiron/sqlx"
	"github.com/samber/do"

	"log"
)

type PostRepositoryImpl struct {
	db *sqlx.DB
}

func NewPostRepository(i *do.Injector) (repository.PostRepository, error) {
	db := do.MustInvoke[*sqlx.DB](i)
	return PostRepositoryImpl{
		db: db,
	}, nil
}

// Create implements repository.PostRepository.
func (p PostRepositoryImpl) Create(ctx context.Context, userID int, title, body string) (int, error) {
	result, err := p.db.ExecContext(ctx,
		`
                INSERT INTO posts
                (
                        posts.user_id,
                        posts.title,
                        posts.body
                )
                VALUES
                (       
                        ?,
                        ?,
                        ?
                )
        `, userID, title, body)
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

// GetDetail implements repository.PostRepository.
func (p PostRepositoryImpl) GetDetail(ctx context.Context, id int) (model.PostDetail, error) {
	post := dao.PostTable{}
	if err := p.db.GetContext(ctx, &post, `
                SELECT
                        posts.id,
                        posts.title,
                        posts.body,
                        posts.user_id,
                        posts.created_at,
                        posts.updated_at,
                        users.name as "user_name"
                FROM
                        posts
                INNER JOIN
                        users
                ON
                        posts.user_id=users.id
                WHERE
                        posts.id=?
                AND
                        posts.deleted_at is null
        `, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.PostDetail{}, nil
		}

		log.Println(err)

		return model.PostDetail{}, err
	}

	return dao.ConvertPostTableToDomainPostDetail(post), nil
}

// List implements repository.PostRepository.
func (p PostRepositoryImpl) List(ctx context.Context) ([]model.Post, error) {
	posts := []dao.PostTable{}
	if err := p.db.SelectContext(ctx, &posts, `
                SELECT
                        posts.id,
                        posts.title,
                        posts.body,
                        posts.user_id,
                        posts.created_at,
                        posts.updated_at,
                        users.name as "user_name"
                FROM
                        posts
                INNER JOIN
                        users
                ON
                        posts.user_id=users.id
                WHERE
                        posts.deleted_at is null
        `); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	log.Println(posts)

	// テーブルのモデルからドメインモデルに変換
	domainPosts := make([]model.Post, len(posts))
	for i, post := range posts {
		domainPosts[i] = dao.ConvertPostTableToDomainPost(post)
	}

	return domainPosts, nil
}

// Update implements repository.PostRepository.
func (p PostRepositoryImpl) Update(ctx context.Context, id int, title, body string) error {
	post := dao.PostTable{}
	if err := p.db.GetContext(ctx, &post, `
                UPDATE
                        posts
                SET
                        title=?,
                        body=?
                WHERE
                        id=?
                        and
                        deleted_at is null
                RETURNING
                        id,
                        title,
                        body,
                        user_id,
                        created_at,
                        updated_at
        `, title, body, id); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// Delete implements repository.PostRepository.
func (p PostRepositoryImpl) Delete(ctx context.Context, id int) error {
	if _, err := p.db.ExecContext(ctx, `
                UPDATE posts
                SET deleted_at=NOW()
                WHERE
                        id=?
                AND
                        deleted_at is null
        `, id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
