package middleware

import (
	"myapp/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine) {
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "It works")
	})
	app.GET("/hello", controllers.HelloWorld)
	app.GET("/posts", controllers.PostsController)
	app.GET("/posts/:postid", controllers.PostDetailController)
	app.GET("/user", func(ctx *gin.Context) {
		ctx.String(200, "User")
	})
}
