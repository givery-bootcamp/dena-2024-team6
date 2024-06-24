package model

type User struct {
	ID       int
	Name     string
	Password string
}

func (u User) IsEmpty() bool {
	return u.ID == 0
}
