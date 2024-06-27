package model

import "time"

type Comment struct {
	ID        int
	PostID    int
	UserID    int
	UserName  string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c Comment) IsEmpty() bool {
	return c.ID == 0
}
