package interfaces

import (
	"myapp/internal/entities"
)

type PostsRepository interface {
	List() ([]*entities.Post, error)
	Get(postID int) (*entities.Post, error)
	CreateComment(postID int, userID int, body string) (*entities.Post, error)
}
