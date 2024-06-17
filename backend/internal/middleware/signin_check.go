package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func SigninCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("token")
		// userがログインしていない場合
		if err != nil {
			ctx.AbortWithError(401, errors.New("unauthorized"))
			return
		}

		secret := []byte("my_secret_key")
		claims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})
		// tokenが不正な場合
		if err != nil {
			ctx.AbortWithError(400, errors.New("invalid token"))
			return
		}

		claimsMap, ok := claims.Claims.(jwt.MapClaims)
		// claimsがMapClaims型でない場合
		if !ok {
			ctx.AbortWithError(500, errors.New("claimsMap type error"))
			return
		}

		userID, ok := claimsMap["ID"].(float64)
		// IDがfloat64型でない場合
		if !ok {
			ctx.AbortWithError(500, errors.New("userID type error"))
			return
		}

		dbCtx, ok := ctx.Get("db")
		if !ok {
			ctx.AbortWithError(500, errors.New("db not found"))
			return
		}
		db, ok := dbCtx.(*gorm.DB)
		if !ok {
			ctx.AbortWithError(500, errors.New("db type error"))
			return
		}
		err = db.Table("users").First(&struct{}{}, "id = ?", int(userID)).Error
		if err != nil {
			ctx.AbortWithError(404, errors.New("user not found"))
			return
		}
		ctx.Next()
	}
}
