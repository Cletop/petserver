package common

import (
	"github.com/gin-gonic/gin"
)

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
