package controllers

import (
	"github.com/gin-gonic/gin"
)

func UserController(ctx *gin.Context) {
	// repository := repositories.NewUserRepository(DB(ctx))
	// usecase := usecases.NewUserUsecase(repository)

	_, err := ctx.Cookie("token")
	if err != nil {
		// handleError(ctx, 401, errors.New("Unauthorized"))
		handleError(ctx, 401, err)
		return
	}

	ctx.String(200, "user")
}
