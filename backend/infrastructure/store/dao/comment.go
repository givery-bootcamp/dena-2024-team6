package dao

import "time"

type CommentTable struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	PostID    int       `db:"post_id"`
	Body      string    `db:"body"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
