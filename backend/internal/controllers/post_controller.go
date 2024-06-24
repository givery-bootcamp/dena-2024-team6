package controllers

import (
	"errors"
	"fmt"
	"myapp/internal/repositories"
	"myapp/internal/usecases"

	"github.com/gin-gonic/gin"
)

type PostParams struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func PostController(ctx *gin.Context) {
	repository := repositories.NewPostsRepository(DB(ctx))
	usecase := usecases.NewPostUsecase(repository)

	var params PostParams
	if err := ctx.ShouldBindJSON(&params); err != nil {
		handleError(ctx, 400, fmt.Errorf("invalid request: %v", err))
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

	post, err := usecase.Execute(userID, params.Title, params.Body)
	if err != nil {
		handleError(ctx, 500, err)
		return
	}

	ctx.JSON(200, post)
}
