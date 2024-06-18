package repositories

import (
	"myapp/internal/entities"

	"gorm.io/gorm"
)

type SigninRepository struct {
	Conn *gorm.DB
}

type User struct {
	ID   int
	Name string
}

func NewSigninRepository(conn *gorm.DB) *SigninRepository {
	return &SigninRepository{
		Conn: conn,
	}
}

func (r *SigninRepository) Signin(username string, password string) (*entities.User, error) {
	var user User
	err := r.Conn.Table("users").Select(`
		id,
		name
	`).Where("BINARY name = ? AND BINARY password = ?", username, password).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &entities.User{
		ID:       user.ID,
		Username: user.Name,
	}, nil
}