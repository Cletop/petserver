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

func VerifyToken(ctx *gin.Context) (bool, uint, string) {
	// get token from cookies
	access_token, refresh_token, cookie_completed := common.GetRenewableCookies(ctx)
	if !cookie_completed {
		ctx.JSON(http.StatusUnauthorized, common.StatusUnauthorizedMessage("Invalid credentials or token."))
		ctx.Abort()
		return false, 0, ""
	}

	// verify tokens
	access_user_id, access_username, access_err := common.VerifyToken(access_token)
	_, _, refresh_err := common.VerifyToken(refresh_token)
	if access_err != nil {
		if refresh_err != nil {
			ctx.JSON(http.StatusUnauthorized, common.StatusUnauthorizedMessage("Invalid credentials	 or token."))
			ctx.Abort()
			return false, 0, ""
		} else {
			// token 续约
			isOk := common.UpdateStorageAuthToken(ctx, access_user_id, access_username)
			if !isOk {
				return false, 0, ""
			}
			ctx.Next()
			return true, access_user_id, access_username
		}
	}

	return true, access_user_id, access_username
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
