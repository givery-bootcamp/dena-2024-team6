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

func (r *UserRepository) Get() (*entities.User, error) {
	return &entities.User{
		ID:       42,
		Username: "Yakisoba Taro",
	}, nil
}
