package middleware

import "github.com/gin-gonic/gin"

func Backlist() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO(lishiwen): implement backlist middleware
		ctx.Next()
	}
}
