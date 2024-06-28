package store

import (
	"context"
	"database/sql"
	"errors"
	"myapp/domain/repository"

	"github.com/jmoiron/sqlx"
	"github.com/samber/do"

	"log"
)

type LikeRepositoryImpl struct {
	db *sqlx.DB
}

func NewLikeRepository(i *do.Injector) (repository.LikeRepository, error) {
	db := do.MustInvoke[*sqlx.DB](i)
	return LikeRepositoryImpl{
		db: db,
	}, nil
}

// Create implements repository.LikeRepository.
func (p LikeRepositoryImpl) Create(ctx context.Context, postID int) (int, error) {
	result, err := p.db.ExecContext(ctx,
		`
                INSERT INTO likes
                (
                    likes.post_id
                )
                VALUES
                (       
                    ?
                )
        `, postID)
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

// Update implements repository.LikeRepository.
func (p LikeRepositoryImpl) Update(ctx context.Context, postID int, value int) error {
	if _, err := p.db.ExecContext(ctx, `
                UPDATE
                    likes
                SET
                    likes=?
                WHERE
                    post_id=?
                AND
                    end_at is null
        `, value, postID); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// Close implements repository.LikeRepository.
func (p LikeRepositoryImpl) Close(ctx context.Context) error {
	if _, err := p.db.ExecContext(ctx, `
		UPDATE
			likes
		SET
			end_at=NOW()
		WHERE
			end_at is null
	`); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Get implements repository.LikeRepository.
func (p LikeRepositoryImpl) Get(ctx context.Context, postID int) (int, error) {
	var likes int
	if err := p.db.GetContext(ctx, &likes, `
		SELECT
			likes.likes
		FROM
			likes
		WHERE
			likes.post_id=?
		AND
			likes.end_at is null
	`, postID); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err)
		}
		return 0, err
	}
	return likes, nil
}
