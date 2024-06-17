package controllers

import (
	"errors"
	"myapp/internal/repositories"
	"myapp/internal/usecases"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func UserController(ctx *gin.Context) {
	repository := repositories.NewUserRepository(DB(ctx))
	usecase := usecases.NewUserUsecase(repository)

	token, err := ctx.Cookie("token")
	if err != nil {
		handleError(ctx, 401, errors.New("unauthorized"))
		return
	}

	secret := []byte("secret")
	claims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		handleError(ctx, 401, errors.New("unauthorized"))
		return
	}

	claimsMap, ok := claims.Claims.(jwt.MapClaims)
	if !ok {
		handleError(ctx, 401, errors.New("unauthorized"))
		return
	}

	userID, ok := claimsMap["ID"].(int)
	if !ok {
		handleError(ctx, 401, errors.New("unauthorized"))
		return
	}

	result, err := usecase.Execute(userID)
	if err != nil {
		handleError(ctx, 500, err)
		return
	}

	ctx.JSON(200, result)
}
