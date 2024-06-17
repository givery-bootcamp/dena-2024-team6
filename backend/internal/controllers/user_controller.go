package controllers

import (
	"errors"
	"fmt"
	"myapp/internal/repositories"
	"myapp/internal/usecases"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func UserController(ctx *gin.Context) {
	repository := repositories.NewUserRepository(DB(ctx))
	usecase := usecases.NewUserUsecase(repository)

	token, err := ctx.Cookie("token")
	// userがログインしていない場合
	if err != nil {
		fmt.Println(err)
		handleError(ctx, 401, errors.New("unauthorized"))
		return
	}

	secret := []byte("my_secret_key")
	claims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	// tokenが不正な場合
	if err != nil {
		handleError(ctx, 400, errors.New("invalid token"))
		return
	}

	claimsMap, ok := claims.Claims.(jwt.MapClaims)
	// claimsがMapClaims型でない場合
	if !ok {
		handleError(ctx, 500, errors.New("claimsMap type error"))
		return
	}

	userID, ok := claimsMap["ID"].(float64)
	// IDがfloat64型でない場合
	if !ok {
		handleError(ctx, 500, errors.New("userID type error"))
		return
	}

	result, err := usecase.Execute(int(userID))
	if err != nil {
		handleError(ctx, 500, err)
		return
	}

	ctx.JSON(200, result)
}
