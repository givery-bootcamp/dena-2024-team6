package dao

import (
	"myapp/domain/model"
	"time"
)

// PostTable は投稿のテーブルを表したモデル
type PostTable struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	UserName  string    `db:"user_name"`
	Title     string    `db:"title"`
	Body      string    `db:"body"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	// deletedAt time.Time `db:"deleted_at"`
}

// ConvertPostTableToDomainPost はテーブルのモデルからドメインモデルに変換する
func ConvertPostTableToDomainPostDetail(pt PostTable) model.PostDetail {
	return model.PostDetail{
		ID:        pt.ID,
		Title:     pt.Title,
		Body:      pt.Body,
		UserID:    pt.UserID,
		UserName:  pt.UserName,
		CreatedAt: pt.CreatedAt,
		UpdatedAt: pt.UpdatedAt,
	}
}

// ConvertPostTableToDomainPost はテーブルのモデルをドメインモデルに変換する
func ConvertPostTableToDomainPost(pt PostTable) model.Post {
	return model.Post{
		ID:        pt.ID,
		Title:     pt.Title,
		UserID:    pt.UserID,
		UserName:  pt.UserName,
		CreatedAt: pt.CreatedAt,
		UpdatedAt: pt.UpdatedAt,
	}
}
