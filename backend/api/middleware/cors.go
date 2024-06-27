package middleware

import (
	"myapp/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	cnf := cors.DefaultConfig()
	cnf.AllowOrigins = []string{config.CorsAllowOrigin}
	cnf.AllowCredentials = true
	return cors.New(cnf)
}
