package usecases

import "myapp/internal/interfaces"

type SigninUsecase struct {
	repository interfaces.SigninRepository
}

func NewSigninUsecase(r interfaces.SigninRepository) *SigninUsecase {
	return &SigninUsecase{
		repository: r,
	}
}

func (u *SigninUsecase) Execute(username string, password string) (string, error) {
	return u.repository.Signin(username, password)
}
