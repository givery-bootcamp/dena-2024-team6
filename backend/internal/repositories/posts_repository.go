package repositories

import (
	"errors"
	"myapp/internal/entities"
	"time"

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
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Comment struct {
	ID        int
	Body      string
	UserId    int
	PostId    int
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
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
	err := r.Conn.Table("posts").Select(`
        posts.id,
        posts.title,
        posts.body,
        posts.user_id,
        users.name as username,
        posts.created_at,
        posts.updated_at
    `).Joins("inner join users on posts.user_id = users.id").Where("posts.id = ?", postID).First(&post).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return convertPostRepositoryModelToEntity(&post), nil
}

func (r *PostsRepository) CreateComment(userID int, postID int, body string) (*entities.Post, error) {
	type CommentComment struct {
		ID     int
		Body   string
		UserId int
		PostID int
	}

	Comment := CommentComment{
		Body:   body,
		UserId: userID,
		PostID: postID,
	}

	err := r.Conn.Table("Comments").Create(&Comment).Error
	if err != nil {
		return nil, err
	}

	return r.Get(Comment.PostID)
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

func convertCommentRepositoryModelToEntity(v *Comment) *entities.Comment {
	return &entities.Comment{
		ID:        v.ID,
		Body:      v.Body,
		UserId:    v.UserId,
		PostID:    v.PostId,
		Username:  v.Username,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
}
