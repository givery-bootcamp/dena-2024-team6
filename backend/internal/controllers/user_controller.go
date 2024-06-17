package controllers

import (
	"errors"
	"fmt"
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func UserController(ctx *gin.Context) {
	repository := repositories.NewUserRepository(DB(ctx))
	usecase := usecases.NewUserUsecase(repository)

	token, err := ctx.Cookie("token")
	if err != nil {
		fmt.Println(err)
		handleError(ctx, 401, errors.New("unauthorized1"))
		return
	}

	secret := []byte("my_secret_key")
	claims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		handleError(ctx, 401, errors.New("unauthorized2"))
		return
	}

	claimsMap, ok := claims.Claims.(jwt.MapClaims)
	if !ok {
		handleError(ctx, 401, errors.New("unauthorized3"))
		return
	}

	userID, ok := claimsMap["ID"].(float64)
	if !ok {
		fmt.Println(reflect.TypeOf(claimsMap["ID"]))
		handleError(ctx, 401, errors.New("unauthorized4"))
		return
	}
	result, err := usecase.Execute(int(userID))
	if err != nil {
		handleError(ctx, 500, err)
		return
	}

	ctx.JSON(200, result)
}
