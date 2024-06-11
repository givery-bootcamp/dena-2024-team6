package controllers

import (
	"errors"
	"myapp/internal/repositories"
	"myapp/internal/usecases"

	"github.com/gin-gonic/gin"
)

type SigninParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SigninController(ctx *gin.Context) {
	repository := repositories.NewSigninRepository(DB(ctx))
	usecase := usecases.NewSigninUsecase(repository)

	var params SigninParams
	if err := ctx.ShouldBindJSON(&params); err != nil {
		handleError(ctx, 400, err)
		return
	}

	if params.Username == "" || params.Password == "" {
		handleError(ctx, 400, errors.New("username and password are required"))
		return
	}

	result, err := usecase.Execute(params.Username, params.Password)
	if err != nil {
		handleError(ctx, 500, err)
		return
	}
	ctx.JSON(200, result)
}
