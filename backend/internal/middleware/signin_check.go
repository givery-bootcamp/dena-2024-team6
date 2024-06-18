package middleware

import (
	"errors"
	"myapp/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type User struct {
	ID   int
	Name string
}

func SigninCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("token")
		// userがログインしていない場合
		if err != nil {
			_ = ctx.AbortWithError(401, errors.New("unauthorized"))
			return
		}

		secret := []byte(config.JwtKey)
		claims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})
		// tokenが不正な場合
		if err != nil {
			_ = ctx.AbortWithError(401, errors.New("invalid token"))
			return
		}

		claimsMap, ok := claims.Claims.(jwt.MapClaims)
		// claimsがMapClaims型でない場合
		if !ok {
			_ = ctx.AbortWithError(500, errors.New("claimsMap type error"))
			return
		}

		userID, ok := claimsMap["ID"].(float64)
		// IDがfloat64型でない場合
		if !ok {
			_ = ctx.AbortWithError(500, errors.New("userID type error"))
			return
		}

		dbCtx, ok := ctx.Get("db")
		if !ok {
			_ = ctx.AbortWithError(500, errors.New("db not found"))
			return
		}
		db, ok := dbCtx.(*gorm.DB)
		if !ok {
			_ = ctx.AbortWithError(500, errors.New("db type error"))
			return
		}
		// ユーザーが存在するか確認
		var user User
		err = db.Table("users").First(&user, "id = ?", int(userID)).Error
		if err != nil {
			_ = ctx.AbortWithError(404, errors.New("user not found"))
			return
		}
		ctx.Set("userID", int(userID))
		ctx.Next()
	}
}
