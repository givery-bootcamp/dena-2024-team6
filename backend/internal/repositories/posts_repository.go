package repositories

import (
	"myapp/internal/entities"

	"gorm.io/gorm"
)

type PostsRepository struct {
	Conn *gorm.DB
}

// This struct is same as entity model
// However define again for training
type Posts struct {
	ID        int
	Title     string
	Body      string
	UserId    int
	Username  string
	CreatedAt string
	UpdatedAt string
}

func NewPostsRepository(conn *gorm.DB) *PostsRepository {
	return &PostsRepository{
		Conn: conn,
	}
}

func (r *PostsRepository) Get() (*entities.Posts, error) {
	obj := Posts{
		ID:        1,
		Title:     "title",
		Body:      "本文",
		UserId:    1,
		Username:  "user1",
		CreatedAt: "2022-05-01T00:00:00Z",
		UpdatedAt: "2022-05-01T00:00:00Z",
	}
	return convertPostsRepositoryModelToEntity(&obj), nil
}

func convertPostsRepositoryModelToEntity(v *Posts) *entities.Posts {
	return &entities.Posts{
		ID:        v.ID,
		Title:     v.Title,
		Body:      v.Body,
		UserId:    v.UserId,
		Username:  v.Username,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
}
