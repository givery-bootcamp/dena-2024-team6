//go:generate mockgen -source=hello_world_repository.go -destination=hello_world_repository_mock.go -package=repository
package repository

import (
	"context"
	"myapp/domain/model"
)

type HelloWorldRepository interface {
	Get(ctx context.Context, lang string) (model.HelloWorld, error)
}
