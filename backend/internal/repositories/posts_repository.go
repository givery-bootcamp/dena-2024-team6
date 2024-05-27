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
type Post struct {
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

func (r *PostsRepository) Get() ([]*entities.Post, error) {
	var posts []Post
	r.Conn.Find(&posts)
	var ent_posts []*entities.Post
	for _, post := range posts {
		ent_posts = append(ent_posts, convertPostRepositoryModelToEntity(&post))
	}
	return ent_posts, nil
}

func convertPostRepositoryModelToEntity(v *Post) *entities.Post {
	return &entities.Post{
		ID:        v.ID,
		Title:     v.Title,
		Body:      v.Body,
		UserId:    v.UserId,
		Username:  v.Username,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
}
