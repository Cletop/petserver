package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func LoadTLS() gin.HandlerFunc {
	return func(c *gin.Context) {
		tls_middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:8080",
		})
		err := tls_middleware.Process(c.Writer, c.Request)
		if err != nil {
			panic(err)
		}
		c.Next()
	}
}
