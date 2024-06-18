package schema

// PpstResponse は投稿のAPIレスポンスモデル「
type PostResponse struct {
	ID    int    `json:"post_id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	UserResponse
}
