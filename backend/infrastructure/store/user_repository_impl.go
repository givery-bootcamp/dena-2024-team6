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

type UserRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserRepositoryImpl(i *do.Injector) (repository.UserRepository, error) {
	db := do.MustInvoke[*sqlx.DB](i)
	return UserRepositoryImpl{
		db: db,
	}, nil
}

// GetByID implements repository.UserRepository.
func (u UserRepositoryImpl) GetByID(ctx context.Context, id int) (model.User, error) {
	user := dao.UserTable{}
	if err := u.db.GetContext(ctx, &user, `
		SELECT
			id,
			name,
			password
		FROM
			users
		WHERE
			id=?
	`, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, nil
		}
		log.Println(err)

		return model.User{}, err
	}

	return dao.ConvertUserTableToDomainUser(user), nil
}

// GetByUserNameAndPassword implements repository.UserRepository.
func (u UserRepositoryImpl) GetByUserNameAndPassword(ctx context.Context, userName string, password string) (model.User, error) {
	user := dao.UserTable{}
	if err := u.db.GetContext(ctx, &user, `
		SELECT
			id,
			name,
			password
		FROM
			users
		WHERE
			name=?
		AND
			password=?
	`, userName, password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, nil
		}
		log.Println(err)

		return model.User{}, err
	}

	return dao.ConvertUserTableToDomainUser(user), nil
}
