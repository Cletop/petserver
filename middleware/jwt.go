package middleware

import (
	"net/http"
	"strconv"

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

// XXX: refactor this code
func VerifyToken(ctx *gin.Context) (bool, uint, string) {
	access_token, _ := common.GetAuthCookie(ctx, common.AccessToken)
	refresh_token, _ := common.GetAuthCookie(ctx, common.RefreshToken)

	// skip verify token if not exist
	if access_token == "" && refresh_token == "" {
		ctx.JSON(http.StatusUnauthorized, common.StatusUnauthorizedMessage("Invalid credentials or token."))
		ctx.Abort()
		return false, 0, ""
	}

	// 存在access_token，验证access_token是否合法
	if access_token != "" {
		// refresh token
		access_token_user_id, access_token_username, access_token_err := common.VerifyToken(access_token)
		if access_token_err == nil {
			ctx.Next()
			return true, access_token_user_id, access_token_username
		} else {
			return false, 0, ""
		}
	} else {
		refresh_token_user_id, refresh_token_username, refresh_token_err := common.VerifyToken(refresh_token)
		if refresh_token_err != nil {
			ctx.JSON(http.StatusUnauthorized, common.StatusUnauthorizedMessage("Invalid credentials or token."))
			ctx.Abort()
			return false, 0, ""
		} else {
			isOK := common.UpdateStorageAuthToken(ctx, refresh_token_user_id, refresh_token_username)
			if !isOK {
				return false, 0, ""
			}
			ctx.Next()
			return true, refresh_token_user_id, refresh_token_username
		}
	}
}

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if NotNeedWTAuth(ctx) {
			ctx.Next()
			return
		}

		passed, user_id, username := VerifyToken(ctx)
		if passed {
			ctx.Set("user_id", user_id)
			ctx.Set("username", username)
			ctx.Next()
		}
	}
}

func RequestedSelfGet(relationKey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// NOTE: 前面的全局过滤器已经验证了Token合法性，这里只需要验证是否是自己的请求
		token, _ := ctx.Cookie(common.AccessToken)
		access_token_user_id, access_token_username, _ := common.VerifyToken(token)

		var err bool = false

		if relationKey == "uid" {
			id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
			if uint(id) != access_token_user_id {
				err = true
			}
		}

		if relationKey == "username" {
			username := ctx.Param("username")
			if username != access_token_username {
				err = true
			}
		}

		if !err {
			ctx.Next()
		} else {
			// 只允许访问自己的信息
			ctx.JSON(
				http.StatusUnauthorized,
				common.StatusRequestedSelfMessage("exceeding access boundaries, allowing access only to your own resources"),
			)
		}
	}
}
