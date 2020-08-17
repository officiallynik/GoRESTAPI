package middlewares

import (
	"github.com/gin-gonic/gin"
)

// AuthenticationMiddleware ==> logic to authentication
func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")

		if len(bearer) != 0 {
			c.Next()
		} else {
			c.JSON(400, "Not Authenticated")
			c.Abort()
		}
	}
}

// AuthorizationMiddleware ==> logic to authorize
func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")
		if bearer == "something" {
			c.Next()
		} else {
			c.JSON(400, "Not Authorized")
			c.Abort()
		}
	}
}
