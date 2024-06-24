package controllers

import "github.com/gin-gonic/gin"

func PostController(ctx *gin.Context) {
	ctx.String(200, "Post created!!!")
}
