package repositories

import (
	"myapp/internal/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{
		Conn: conn,
	}
}

func (r *UserRepository) Get(id int) (*entities.User, error) {
	var user User
	err := r.Conn.Table("users").Select(`
		id,
		name
	`).Where("id = ?", id).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &entities.User{
		ID:       user.ID,
		Username: user.Name,
	}, nil
}
