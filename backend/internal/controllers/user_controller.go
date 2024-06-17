package controllers

import "github.com/gin-gonic/gin"

func UserController(ctx *gin.Context) {
	ctx.String(200, "User")
}
