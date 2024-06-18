package controller

import (
	"context"
	"myapp/api/schema"
	"myapp/application"
	"myapp/config"
	"myapp/domain/apperror"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type PostController struct {
	listPostUsecase application.ListPostUsecase
	getPostUsecase  application.GetPostDetailUsecase
}

func NewPostController(i *do.Injector) (*PostController, error) {
	listPostUsecase := do.MustInvoke[application.ListPostUsecase](i)
	getPostUsecase := do.MustInvoke[application.GetPostDetailUsecase](i)
	return &PostController{
		listPostUsecase: listPostUsecase,
		getPostUsecase:  getPostUsecase,
	}, nil
}

func (pc PostController) ListPost(ctx *gin.Context) {
	pc.listPostUsecase.Execute(ctx)
	// TODO: 一旦仮のデータを返す
	ctx.JSON(200, []schema.PostResponse{
		{
			ID:    1,
			Title: "hoge",
			Body:  "hogehoge",
			UserResponse: schema.UserResponse{
				ID:       2,
				UserName: "fugafuga",
			},
		},
		{
			ID:    2,
			Title: "hoge",
			Body:  "hogehoge",
			UserResponse: schema.UserResponse{
				ID:       2,
				UserName: "fugafuga",
			},
		},
	})
}

func (pc PostController) GetPost(c *gin.Context) {
	ctx, cancel := context.WithDeadline(c, time.Now().Add(time.Duration(config.DefaultTimeoutSecond)*time.Second))
	defer cancel()

	postID, err := strconv.Atoi(c.Param("postid"))
	if err != nil {
		c.JSON(400, schema.NewErrorResponse(
			apperror.New(apperror.CodeInvalidArgument, "invalid argument"),
		))
	}

	result, err := pc.getPostUsecase.Execute(ctx, application.GetPostDetailUsecaseInput{
		ID: postID,
	})
	if apperror.Is(err, apperror.CodeNotFound) {
		c.JSON(404, schema.NewErrorResponse(err))
		return
	}
	if apperror.Is(err, apperror.CodeInternalServer) {
		c.JSON(500, schema.NewErrorResponse(err))
		return
	}
	// TODO: 一旦仮のデータを返す
	c.JSON(200, schema.PostResponse{
		ID:    result.Post.ID,
		Title: result.Post.Title,
		Body:  result.Post.Body,
		UserResponse: schema.UserResponse{
			ID:       result.Post.UserID,
			UserName: result.Post.UserName,
		},
	})
}
