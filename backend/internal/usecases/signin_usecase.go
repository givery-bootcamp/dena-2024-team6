package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/interfaces"
	"time"

	"github.com/golang-jwt/jwt"
)

type SigninUsecase struct {
	repository interfaces.SigninRepository
}

func NewSigninUsecase(r interfaces.SigninRepository) *SigninUsecase {
	return &SigninUsecase{
		repository: r,
	}
}

func (u *SigninUsecase) Execute(username string, password string) (*entities.User, *jwt.Token, error) {
	result, err := u.repository.Signin(username, password)
	if err != nil {
		return nil, nil, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":        result.ID,
		"ExpiresAt": time.Now().Add(time.Hour * 24).Unix(),
	})
	return result, token, nil
}
