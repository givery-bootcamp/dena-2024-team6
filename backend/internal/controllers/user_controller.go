package controllers

import (
	"errors"
	"myapp/internal/repositories"
	"myapp/internal/usecases"

	"github.com/gin-gonic/gin"
)

func UserController(ctx *gin.Context) {
	repository := repositories.NewUserRepository(DB(ctx))
	usecase := usecases.NewUserUsecase(repository)

	userIDCtx, ok := ctx.Get("userID")
	if !ok {
		handleError(ctx, 500, errors.New("user not found"))
		return
	}
	userID, ok := userIDCtx.(int)
	if !ok {
		handleError(ctx, 500, errors.New("userID type error"))
		return
	}

	result, err := usecase.Execute(userID)
	if err != nil {
		handleError(ctx, 500, err)
		return
	}
	ctx.JSON(200, result)
}
