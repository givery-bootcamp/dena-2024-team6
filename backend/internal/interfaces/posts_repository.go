package interfaces

import (
	"myapp/internal/entities"
)

type PostsRepository interface {
	Get() (*entities.Posts, error)
}
