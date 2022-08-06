package common

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const AccessToken string = "access_token"
const RefreshToken string = "refresh_token"

func SetHttpOnlyCookie(ctx *gin.Context, key, value string, maxAge int) {
	// Set HttpOnly cookie

	// only allow cross-origin requests from the same origin to prevent CSRF attacks
	domain := strings.Split(strings.Split(ctx.Request.Header.Get("Origin"), ":")[1], "//")[1] // real domain
	// host := ctx.Request.Host                                          // hostname ,proxy server hostname

	/**
	* key: cookie name
	* value: cookie value
	* maxAge: cookie max age in seconds
	* path: cookie path
	* domain: cookie domain
	* secure: cookie secure flag (only HTTPS)
	* httpOnly: cookie http only flag (Do not allow javascript access/modification)
	 */
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(key, value, maxAge, "/", domain, false, true)
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
	SetHttpOnlyCookie(ctx, AccessToken, value, 10*60) // 10 minutes
}
func SetRefreshCookie(ctx *gin.Context, value string) {
	SetSecureCookie(ctx, RefreshToken, value, 7*60*60*24) // 7 days
}

func GetAuthCookie(ctx *gin.Context, key string) (string, error) {
	// Get cookie value
	cookie, err := ctx.Cookie(key)
	return cookie, err
}

func GetRenewableCookies(ctx *gin.Context) (string, string, bool) {
	// Get renewable cookies
	access_token, access_token_err := GetAuthCookie(ctx, AccessToken)
	refresh_token, refresh_token_err := GetAuthCookie(ctx, RefreshToken)
	if access_token_err != nil || refresh_token_err != nil {
		return "", "", false
	}
	return access_token, refresh_token, true
}

func UpdateStorageAuthToken(ctx *gin.Context, user_id uint, username string) bool {
	access_token, refresh_token, err := CreateRenewableToken(user_id, username)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			StatusInternalServerErrorMessage("failed to create token"),
		)
		return false
	}

	SetAuthCookie(ctx, access_token)
	SetRefreshCookie(ctx, refresh_token)

	return true
}

func DeleteStorageAuthToken(ctx *gin.Context) {
	SetHttpOnlyCookie(ctx, AccessToken, "", -1)
	SetHttpOnlyCookie(ctx, RefreshToken, "", -1)
}
