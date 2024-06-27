package middleware

import (
	"github.com/gin-gonic/gin"
)

// TODO: 新しい実装に置き換える
func Transaction() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// db := external.DB.Begin()
		// defer func() {
		// 	if 400 <= ctx.Writer.Status() {
		// 		db.Rollback()
		// 		return
		// 	}
		// 	db.Commit()
		// }()
		// ctx.Set("db", db)
		// ctx.Next()
	}
}
