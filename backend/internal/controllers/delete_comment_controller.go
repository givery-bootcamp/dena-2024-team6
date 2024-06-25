package controllers

import (
	"errors"
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteCommentController(ctx *gin.Context) {
	repository := repositories.NewPostsRepository(DB(ctx))
	usecase := usecases.NewDeleteCommentUsecase(repository)

	CommentID, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		handleError(ctx, 400, err)
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

	err = usecase.Execute(userID, CommentID)
	if err != nil {
		handleError(ctx, 400, err)
		return
	}

	ctx.JSON(204, nil)
}
