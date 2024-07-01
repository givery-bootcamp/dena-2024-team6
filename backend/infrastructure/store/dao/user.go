package dao

import (
	"myapp/domain/model"
)

// UserTable はユーザのテーブルを表したモデル
type UserTable struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
	IconURL  string `db:"icon_url"`
	// createdAt time.Time `db:"created_at"`
	// updatedAt time.Time `db:"updated_at"`
}

// ConvertUserTableToDomainUser はテーブルのモデルからドメインモデルに変換する
func ConvertUserTableToDomainUser(ut UserTable) model.User {
	return model.User{
		ID:       ut.ID,
		Name:     ut.Name,
		Password: ut.Password,
		IconURL:  ut.IconURL,
	}
}
