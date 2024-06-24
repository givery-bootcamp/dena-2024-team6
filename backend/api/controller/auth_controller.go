package controller

import (
	"context"
	"errors"
	"myapp/api/schema"
	"myapp/application"
	"myapp/config"
	"myapp/domain/apperror"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type AuthController struct {
	signinUsecase  application.SigninUsecase
	getUserUsecase application.GetUserUsecase
}

func NewAuthController(i *do.Injector) (*AuthController, error) {
	signinUsecase := do.MustInvoke[application.SigninUsecase](i)
	getUserUsecase := do.MustInvoke[application.GetUserUsecase](i)
	return &AuthController{
		signinUsecase:  signinUsecase,
		getUserUsecase: getUserUsecase,
	}, nil
}

func (ac AuthController) SignIn(c *gin.Context) {
	ctx, cancel := context.WithDeadline(c, time.Now().Add(time.Duration(config.DefaultTimeoutSecond)*time.Second))
	defer cancel()

	var params schema.LoginRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(400, schema.NewErrorResponse(errors.New("invalid argument")))
		return
	}

	result, err := ac.signinUsecase.Execute(ctx, application.SigninUsecaseInput{
		UserName: params.UserName,
		Password: params.Password,
	})
	if apperror.Is(err, apperror.CodeUnauthorized) {
		c.JSON(401, schema.NewErrorResponse(err))
		return
	}
	if apperror.Is(err, apperror.CodeInternalServer) {
		c.JSON(500, schema.NewErrorResponse(err))
		return
	}
	if err != nil {
		c.JSON(500, schema.NewErrorResponse(err))
		return
	}

	c.SetCookie("token", result.Token, 3600*24, "/", config.HostName, false, true)

	c.JSON(200, schema.UserResponse{
		ID:       result.User.ID,
		UserName: result.User.Name,
	})
}

func (ac AuthController) SignOut(ctx *gin.Context) {
	// TODO: あとで実装する
}

func (ac AuthController) GetCurrentUser(c *gin.Context) {
	ctx, cancel := context.WithDeadline(c, time.Now().Add(time.Duration(config.DefaultTimeoutSecond)*time.Second))
	defer cancel()

	userIDAny, ok := c.Get("userID")
	if !ok {
		c.JSON(401, "unauthorized")
	}
	userID, ok := userIDAny.(int)
	if !ok {
		c.JSON(500, "failed to login")
	}

	result, err := ac.getUserUsecase.Execute(ctx, application.GetUserUsecaseInput{
		ID: userID,
	})
	if apperror.Is(err, apperror.CodeNotFound) {
		c.JSON(404, schema.NewErrorResponse(err))
		return
	}
	if apperror.Is(err, apperror.CodeInternalServer) {
		c.JSON(500, schema.NewErrorResponse(err))
		return
	}
	if err != nil {
		c.JSON(500, schema.NewErrorResponse(err))
		return
	}

	c.JSON(200, schema.UserResponse{
		ID:       result.User.ID,
		UserName: result.User.Name,
	})
}