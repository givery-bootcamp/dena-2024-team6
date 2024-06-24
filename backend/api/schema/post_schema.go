package schema

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
