package interfaces

import (
	"myapp/internal/entities"
)

type PostsRepository interface {
	List() ([]*entities.Post, error)
	Get(postID int) (*entities.Post, error)
	Create(userID int, title string, body string) (*entities.Post, error)
	Delete(userID int, postID int) error
	Update(userID int, postID int, title string, body string) (*entities.Post, error)
}
