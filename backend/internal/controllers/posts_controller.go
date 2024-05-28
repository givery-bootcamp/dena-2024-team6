package controllers

import (
	"errors"
	"myapp/internal/repositories"
	"myapp/internal/usecases"

	"github.com/gin-gonic/gin"
)

func PostsController(ctx *gin.Context) {
	repository := repositories.NewPostsRepository(DB(ctx))
	usecase := usecases.NewPostsUsecase(repository)

	result, err := usecase.Execute()
	if err != nil {
		handleError(ctx, 500, err)
		return
	}
	if len(result) == 0 {
		handleError(ctx, 404, errors.New("Not found"))
		return
	}

	ctx.JSON(200, result)
}
