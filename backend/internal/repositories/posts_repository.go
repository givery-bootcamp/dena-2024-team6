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

func (r *PostsRepository) Delete(UserId int, postID int) error {
	//削除する投稿がログインユーザーのものか確認
	deletetarget := Post{}
	err := r.Conn.Table("posts").Where("id = ?", postID).First(&deletetarget).Error
	if err != nil {
		return err
	}
	targetUserID := deletetarget.UserId
	if targetUserID != UserId {
		return errors.New("this post is not yours")
	}
	err = r.Conn.Table("posts").Where("id = ?", postID).Delete(&Post{}).Error
	if err != nil {
		return err
	}
	return nil
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
