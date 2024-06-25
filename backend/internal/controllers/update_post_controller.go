package controllers

import (
	"errors"
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdatePostController(ctx *gin.Context) {
	repository := repositories.NewPostsRepository(DB(ctx))
	usecase := usecases.NewUpdatePostUsecase(repository)

	postID, err := strconv.Atoi(ctx.Param("postid"))
	if err != nil {
		handleError(ctx, 400, err)
		return
	}

	var params PostParams
	if err := ctx.ShouldBindJSON(&params); err != nil {
		handleError(ctx, 400, errors.New("invalid request"))
		return
	}

	if params.Title == "" || params.Body == "" {
		handleError(ctx, 400, errors.New("title and body are required"))
		return
	}

	if len(params.Title) > 100 {
		handleError(ctx, 400, errors.New("title must be less than 100 characters"))
		return
	}

	rawUserID, ok := ctx.Get("userID")
	if !ok {
		handleError(ctx, 500, errors.New("userID not found"))
		return
	}

	userID, ok := rawUserID.(int)
	if !ok {
		handleError(ctx, 500, errors.New("userID is invalid"))
		return
	}

	post, err := usecase.Execute(userID, postID, params.Title, params.Body)
	if err != nil {
		handleError(ctx, 500, err)
		return
	}

	ctx.JSON(200, post)
}
