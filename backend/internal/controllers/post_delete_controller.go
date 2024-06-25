package controllers

import (
	"errors"
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeletePostController(ctx *gin.Context) {
	repository := repositories.NewPostsRepository(DB(ctx))
	usecase := usecases.NewPostDeleteUsecase(repository)

	rawPostID := ctx.Param("postid")
	postID, _err := strconv.Atoi(rawPostID)
	if _err != nil {
		handleError(ctx, 400, _err)
		return
	}
	rawUserID, ok := ctx.Get("userID")
	if !ok {
		handleError(ctx, 500, errors.New("userID not found"))
		return
	}

	userID, ok := rawUserID.(int)
	if !ok {
		handleError(ctx, 500, errors.New("unauthorized"))
		return
	}

	err := usecase.Execute(userID, postID)
	if err != nil {
		handleError(ctx, 400, err)
		return
	}
	ctx.Status(204)

}
