package dao

import (
	"myapp/domain/model"
	"time"
)

type CommentTable struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	UserName  string    `db:"user_name"`
	PostID    int       `db:"post_id"`
	Body      string    `db:"body"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func ConvertCommentTableToDomainComment(ct CommentTable) model.Comment {
	return model.Comment{
		ID:        ct.ID,
		PostID:    ct.PostID,
		UserID:    ct.UserID,
		UserName:  ct.UserName,
		Body:      ct.Body,
		CreatedAt: ct.CreatedAt,
		UpdatedAt: ct.UpdatedAt,
	}
}
