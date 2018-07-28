package middleware

import (
	"../../handle"
	"../../pkg/errno"
	"../../pkg/token"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		if _, err := token.ParseRequest(c); err != nil {
			handle.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Next()
	}

}
