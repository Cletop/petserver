package middleware

import (
	"net/http"

	"github.com/chagspace/petserver/common"
	"github.com/gin-gonic/gin"
)

// not need auth list
var NotNeedWTAuthList = []string{
	"/api/v1/users",
	"/api/v1/users/login",
}

func NotNeedWTAuth(ctx *gin.Context) bool {
	method := ctx.Request.Method == "POST"
	path := ctx.Request.URL.Path // /api/v1/users

	// XXX: refactor this code
	for _, v := range NotNeedWTAuthList {
		if v == path && method {
			return true
		}
	}

	return false
}

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		notNeedWTAuth := NotNeedWTAuth(ctx)

		if notNeedWTAuth {
			ctx.Next()
			return
		}

		value, err := common.GetAuthCookie(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":   1,
				"status": "unauthorized",
				"msg":    err.Error(),
			})
			ctx.Abort()
			return
		}

		// verify tokens
		user_id, username, err := common.VerifyToken(value)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":   1,
				"status": "unauthorized",
				"msg":    err.Error(),
			})
			ctx.Abort()
			return
		}

		ctx.Set("user_id", user_id)
		ctx.Set("username", username)

		ctx.Next()
	}
}
