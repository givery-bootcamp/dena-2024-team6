package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/interfaces"
)

type SigninUsecase struct {
	repository interfaces.SigninRepository
}

func NewSigninUsecase(r interfaces.SigninRepository) *SigninUsecase {
	return &SigninUsecase{
		repository: r,
	}
}

func (u *SigninUsecase) Execute(username string, password string) (*entities.User, error) {
	return u.repository.Signin(username, password)
}
