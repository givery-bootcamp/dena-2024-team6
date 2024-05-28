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

func (r *PostsRepository) List() ([]*entities.Post, error) {
	var posts []Post
	if err := r.Conn.Table("posts").Select(`
        posts.id,
        posts.title,
        posts.body,
        posts.user_id,
        users.name as username,
        posts.created_at,
        posts.updated_at
    `).Joins("inner join users on posts.user_id = users.id").Order("posts.id DESC").Scan(&posts).Error; err != nil {
		return nil, err
	}
	var ent_posts []*entities.Post
	for _, post := range posts {
		ent_posts = append(ent_posts, convertPostRepositoryModelToEntity(&post))
	}
	return ent_posts, nil
}

func (r *PostsRepository) Get(postID int) (*entities.Post, error) {
	var post Post
	if err := r.Conn.Table("posts").Select(`
        posts.id,
        posts.title,
        posts.body,
        posts.user_id,
        users.name as username,
        posts.created_at,
        posts.updated_at
    `).Joins("inner join users on posts.user_id = users.id").Where("posts.id = ?", postID).Scan(&post).Error; err != nil {
		return nil, err
	}

	return convertPostRepositoryModelToEntity(&post), nil
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
