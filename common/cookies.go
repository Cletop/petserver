package common

import (
	"github.com/gin-gonic/gin"
)

var AccessToken string = "access_token"
var AuthTokenMaxAge int = 7 * 60 * 60 * 24

func SetHttpOnlyCookie(ctx *gin.Context, key, value string, maxAge int) {
	// Set HttpOnly cookie

	// only allow cross-origin requests from the same origin to prevent CSRF attacks
	origin := ctx.Request.Header.Get("Origin")

	/**
	* key: cookie name
	* value: cookie value
	* maxAge: cookie max age in seconds
	* path: cookie path
	* domain: cookie domain
	* secure: cookie secure flag (only HTTPS)
	* httpOnly: cookie http only flag (Do not allow javascript access/modification)
	 */
	ctx.SetCookie(key, value, maxAge, "/", origin, false, true)
}

func SetSecureCookie(ctx *gin.Context, key, value string, maxAge int) {
	// Set Secure cookie

	// only allow cross-origin requests from the same origin to prevent CSRF attacks
	origin := ctx.Request.Header.Get("Origin")

	// Set secure cookies
	// only allow at https requests and not allow javascript access/modification
	ctx.SetCookie(key, value, maxAge, "/", origin, true, true)
}

func SetAuthCookie(ctx *gin.Context, value string) {
	// Set Auth cookie
	SetSecureCookie(ctx, AccessToken, value, AuthTokenMaxAge)
}

func GetAuthCookie(ctx *gin.Context) (string, error) {
	// Get cookie value
	cookie, err := ctx.Cookie(AccessToken)
	return cookie, err
}
