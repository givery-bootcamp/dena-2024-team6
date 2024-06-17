package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/interfaces"
)

type UserUsecase struct {
	repository interfaces.UserRepository
}

func NewUserUsecase(r interfaces.UserRepository) *UserUsecase {
	return &UserUsecase{
		repository: r,
	}
}

func (u *UserUsecase) Execute() (*entities.User, error) {
	return u.repository.Get()
}
