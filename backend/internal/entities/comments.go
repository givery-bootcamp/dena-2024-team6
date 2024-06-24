package entities

import "time"

type Comment struct {
	ID        int       `json:"id"`
	Body      string    `json:"body"`
	UserId    int       `json:"user_id"`
	PostID    int       `json:"post_id"`
	Username  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
