package controllers

import (
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostDetailController(ctx *gin.Context) {
	repository := repositories.NewPostsRepository(DB(ctx))
	usecase := usecases.NewPostDetailUsecase(repository)

	postID, err := strconv.Atoi(ctx.Param("postid"))
	if err != nil {
		handleError(ctx, 500, err)
		return
	}

	result, err := usecase.Execute(postID)
	if err != nil {
		handleError(ctx, 500, err)
		return
	}

	ctx.JSON(200, result)
}
