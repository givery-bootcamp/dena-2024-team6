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
	DeletedAt gorm.DeletedAt
}

type Comment struct {
	ID        int
	Body      string
	UserId    int
	PostId    int
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
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


func (r *PostsRepository) Create(userID int, title string, body string) (*entities.Post, error) {
	type PostPost struct {
		ID     int
		Title  string
		Body   string
		UserId int
	}

	post := PostPost{
		Title:  title,
		Body:   body,
		UserId: userID,
	}

	err := r.Conn.Table("posts").Create(&post).Error
	if err != nil {
		return nil, err
	}

	return r.Get(post.ID)
}

func (r *PostsRepository) Update(userID int, postID int, title string, body string) (*entities.Post, error) {
	oldPost, err := r.Get(postID)
	if err != nil {
		return nil, err
	}
	if oldPost == nil {
		return nil, errors.New("not found")
	}
	if oldPost.UserId != userID {
		return nil, errors.New("unauthorized")
	}

	type PostPut struct {
		Title string
		Body  string
	}

	post := PostPut{
		Title: title,
		Body:  body,
	}

	err = r.Conn.Table("posts").Where("id = ?", postID).Updates(&post).Error
	if err != nil {
		return nil, err
	}
	return r.Get(postID)
}


func (r *PostsRepository) CreateComment(userID int, postID int, body string) (*entities.Post, error) {
	type PostComment struct {
		ID     int
		Body   string
		UserId int
		PostID int
	}

	Comment := PostComment{
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

func (r *PostsRepository) UpdateComment(userID int, CommentID int, postID int, body string) (*entities.Post, error) {
	oldComment, err := r.Get(postID)
	if err != nil {
		return nil, err
	}
	if oldComment == nil {
		return nil, errors.New("not found")
	}
	if oldComment.UserId != userID {
		return nil, errors.New("unauthorized")
	}

	type PostComment struct {
		ID     int
		Body   string
		UserId int
		PostID int
	}

	Comment := PostComment{
		Body:   body,
		UserId: userID,
		PostID: postID,
	}

	err = r.Conn.Table("Comments").Where("id = ?", CommentID).Updates(&Comment).Error
	if err != nil {
		return nil, err
	}
	return r.Get(postID)
}

func (r *PostsRepository) DeleteComment(userID int, CommentID int) error {
	var comment Comment
	err := r.Conn.Table("Comments").Where("id = ?", CommentID).First(&comment).Error
	if err != nil {
		return errors.New("not found")
	}
	if comment.UserId != userID {
		return errors.New("unauthorized")
	}

	err = r.Conn.Table("Comments").Where("id = ?", CommentID).Delete(&Comment{}).Error
	if err != nil {
		return err
	}
	return nil
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
