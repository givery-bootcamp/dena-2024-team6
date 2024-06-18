package schema

// UserResponse はユーザのAPIレスポンスモデル
type UserResponse struct {
	ID       int    `json:"user_id"`
	UserName string `json:"user_name"`
}

type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
