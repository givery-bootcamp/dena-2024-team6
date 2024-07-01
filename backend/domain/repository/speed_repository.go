package repository

import (
	"context"
	"myapp/domain/model"
)

type SpeedRepository interface {
	List(ctx context.Context) ([]model.Speed, error)
}
