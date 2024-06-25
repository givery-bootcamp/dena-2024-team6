package schema

import "time"

type PostRequest struct {
	ID string `path:"id" example:"1"`
}

// PpstResponse は投稿のAPIレスポンスモデル「
type PostResponse struct {
	ID    int    `json:"post_id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	UserResponse
}

type CommentListRequest struct {
	PostID string `path:"postId" example:"1"`
}

type CreateCommentRequest struct {
	PostID string `path:"postId" example:"1"`
	Body   string `json:"body"`
}

type CommentResponse struct {
	ID     int    `json:"id"`
	PostID int    `json:"post_id"`
	Body   string `json:"body"`
	UserResponse
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
