package controllers

import (
	"github.com/gin-gonic/gin"
)

func SignoutController(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
	ctx.Status(200)
}
