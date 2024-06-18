package controller

import (
	"myapp/api/schema"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type AuthController struct{}

func NewAuthController(i *do.Injector) (*AuthController, error) {
	return &AuthController{}, nil
}

func (ac AuthController) SignIn(ctx *gin.Context) {
	// TODO: 一旦仮のデータを返す
	ctx.JSON(200, schema.UserResponse{
		ID:       2,
		UserName: "fugafuga",
	})
}

func (ac AuthController) SignOut(ctx *gin.Context) {
	// TODO: あとで実装する
}

func (ac AuthController) GetCurrentUser(ctx *gin.Context) {
	// TODO: 一旦仮のデータを返す
	ctx.JSON(200, schema.UserResponse{
		ID:       2,
		UserName: "fugafuga",
	})
}
