package model

type User struct {
	ID       int
	Name     string
	Password string
	IconURL  string
}

func (u User) IsEmpty() bool {
	return u.ID == 0
}
