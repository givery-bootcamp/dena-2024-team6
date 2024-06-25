package controllers

import (
	"errors"
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateCommentController(ctx *gin.Context) {
	repository := repositories.NewPostsRepository(DB(ctx))
	usecase := usecases.NewUpdateCommentUsecase(repository)

	CommentID, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		handleError(ctx, 400, err)
		return
	}

	var params CommentParams
	if err := ctx.ShouldBindJSON(&params); err != nil {
		handleError(ctx, 400, errors.New("invalid request"))
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

	Comment, err := usecase.Execute(userID, CommentID, intPostID, params.Body)
	if err != nil {
		handleError(ctx, 400, err)
		return
	}

	ctx.JSON(200, Comment)
}
