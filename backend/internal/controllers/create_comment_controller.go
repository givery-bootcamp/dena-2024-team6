package controllers

import (
	"errors"
	"fmt"
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentParams struct {
	PostID string `json:"post_id"`
	Body   string `json:"body"`
}

func CreateCommentController(ctx *gin.Context) {
	repository := repositories.NewPostsRepository(DB(ctx))
	usecase := usecases.NewCommentUsecase(repository)

	var params CommentParams
	if err := ctx.ShouldBindJSON(&params); err != nil {
		handleError(ctx, 400, fmt.Errorf("invalid request: %v", err))
		return
	}

	if params.PostID == "" || params.Body == "" {
		handleError(ctx, 400, errors.New("title and body are required"))
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

	intPostID, err := strconv.Atoi(params.PostID)
	if err != nil {
		handleError(ctx, 400, errors.New("post_id must be a number"))
		return
	}

	post, err := usecase.Execute(userID, intPostID, params.Body)
	if err != nil {
		handleError(ctx, 500, err)
		return
	}

	ctx.JSON(200, post)
}
