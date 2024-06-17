package usecases

import (
	"fmt"
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
	var jwtKey = []byte("my_secret_key")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":        result.ID,
		"ExpiresAt": time.Now().Add(time.Hour * 24).Unix(),
	})
	token.Raw, _ = token.SignedString(jwtKey)

	fmt.Printf("result.ID: %v\n", result.ID)
	fmt.Printf("ExpiresAt: %v\n", time.Now().Add(time.Hour*24).Unix())
	fmt.Printf("tokenString: %v\n", token.Raw)
	fmt.Printf("Generate Token with claims: %v\n", token.Claims)
	return result, token, nil
}
