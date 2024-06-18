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
)

type HelloWorldRepositoryImpl struct {
	db *sqlx.DB
}

func NewHelloWorldRepository(i *do.Injector) (repository.HelloWorldRepository, error) {
	db := do.MustInvoke[*sqlx.DB](i)
	return HelloWorldRepositoryImpl{
		db: db,
	}, nil
}

// Get implements repository.HelloWorldRepository.
func (h HelloWorldRepositoryImpl) Get(ctx context.Context, lang string) (model.HelloWorld, error) {
	helloWorld := dao.HelloWorldTable{}
	if err := h.db.GetContext(ctx, &helloWorld, `
		SELECT
			lang,
			message
		FROM
			hello_worlds
		WHERE
			lang=?
	`, lang); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.HelloWorld{}, nil
		}

		return model.HelloWorld{}, err
	}

	return dao.ConvertHelloWorldTableToDomainHelloWorld(helloWorld), nil
}
