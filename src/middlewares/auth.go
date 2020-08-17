package middlewares

import (
	"github.com/gin-gonic/gin"
)

// AuthMiddleware ==> logic to authentication
func AuthMiddleware() gin.HandlerFunc {
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
