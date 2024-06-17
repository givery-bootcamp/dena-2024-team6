package model

import "time"

// Post は投稿 (body抜き)
// | 主に投稿一覧で使用するドメインモデルで、Bodyまで取得したい場合は PostDetail を使う
type Post struct {
	ID        int
	Title     string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// PostDetail は投稿の詳細
// | Bodyは長文のため、詳細の時にしか取得しないようにしている
type PostDetail struct {
	ID        int
	Title     string
	Body      string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
