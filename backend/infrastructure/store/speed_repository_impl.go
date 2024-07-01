package store

import (
	"context"
	"errors"

	"myapp/domain/model"
	"myapp/domain/repository"

	"github.com/jmoiron/sqlx"
	"github.com/samber/do"
)

type SpeedRepositoryImpl struct {
	db *sqlx.DB
}

func NewSpeedRepositoryImpl(i *do.Injector) (repository.SpeedRepository, error) {
	db := do.MustInvoke[*sqlx.DB](i)
	return SpeedRepositoryImpl{
		db: db,
	}, nil
}

func (s SpeedRepositoryImpl) List(ctx context.Context) ([]model.Speed, error) {
	return nil, errors.New("not implemented")
}
