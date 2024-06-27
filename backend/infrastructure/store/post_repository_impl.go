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
func (p PostRepositoryImpl) Create(ctx context.Context, userID int, title, body string) (model.Post, error) {
	post := dao.PostTable{}
	if err := p.db.GetContext(ctx, &post, `
        INSERT INTO
                posts (user_id, title, body)
        VALUES
        (
                ?,
                ?,
                ?
        )
        RETURNING
                id,
                title,
                body,
                user_id,
                created_at,
                updated_at
        `, userID, title, body); err != nil {
		log.Println(err)
		return model.Post{}, err
	}

	return dao.ConvertPostTableToDomainPost(post), nil
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
        `); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	// テーブルのモデルからドメインモデルに変換
	domainPosts := make([]model.Post, len(posts))
	for i, post := range posts {
		domainPosts[i] = dao.ConvertPostTableToDomainPost(post)
	}

	return domainPosts, nil
}
