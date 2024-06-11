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
	app.POST("/signin", func(ctx *gin.Context) {
		// Get two parameters from JSON: username (string) and password (string)

		var json struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := ctx.ShouldBindJSON(&json); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"username": json.Username, "password": json.Password})
	})
}
