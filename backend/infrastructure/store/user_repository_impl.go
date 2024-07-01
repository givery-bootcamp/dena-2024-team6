package store

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
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

// Create implements repository.UserRepository.
func (u UserRepositoryImpl) Create(ctx context.Context, userName, password string) (model.User, error) {
	hashedPassword := sha256.Sum256([]byte(password))
	strPassword := hex.EncodeToString(hashedPassword[:])

	result, err := u.db.ExecContext(ctx, `
                INSERT INTO
                        users (name, password)
                VALUES
                        (?, ?)
        `,
		userName,
		strPassword,
	)
	if err != nil {
		log.Println(err)
		return model.User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return model.User{}, err
	}

	return model.User{
		ID:       int(id),
		Name:     userName,
		Password: strPassword,
	}, nil
}

// Exists implements repository.UserRepository.
func (u UserRepositoryImpl) Exists(ctx context.Context, userName string) (bool, error) {
	var count int
	if err := u.db.GetContext(ctx, &count, `
                SELECT
                        COUNT(*)
                FROM
                        users
                WHERE
                        BINARY name=?
        `, userName); err != nil {
		log.Println(err)
		return false, err
	}

	return count > 0, nil
}

// GetByID implements repository.UserRepository.
func (u UserRepositoryImpl) GetByID(ctx context.Context, id int) (model.User, error) {
	user := dao.UserTable{}
	if err := u.db.GetContext(ctx, &user, `
		SELECT
			id,
			name,
			password,
                        icon_url
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
	hashedPassword := sha256.Sum256([]byte(password))
	strPassword := hex.EncodeToString(hashedPassword[:])

	user := dao.UserTable{}
	if err := u.db.GetContext(ctx, &user, `
		SELECT
			id,
			name,
			password,
                        icon_url
		FROM
			users
		WHERE
			BINARY name=?
		AND
			BINARY password=?
	`, userName, strPassword); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, nil
		}
		log.Println(err)

		return model.User{}, err
	}

	return dao.ConvertUserTableToDomainUser(user), nil
}
