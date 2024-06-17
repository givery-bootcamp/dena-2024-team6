package controller

import (
	"myapp/cmd/api/schema"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type PostController struct{}

func NewPostController(i *do.Injector) (*PostController, error) {
	return &PostController{}, nil
}

func (pc PostController) ListPost(ctx *gin.Context) {
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

func (pc PostController) GetPost(ctx *gin.Context) {
	// TODO: 一旦仮のデータを返す
	ctx.JSON(200, schema.PostResponse{
		ID:    1,
		Title: "hoge",
		Body:  "hogehoge",
		UserResponse: schema.UserResponse{
			ID:       2,
			UserName: "fugafuga",
		},
	})
}
