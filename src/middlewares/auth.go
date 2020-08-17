package middlewares

import (
	"blogapp/src/routes"

	"github.com/gin-gonic/gin"
)

// AuthenticationMiddleware ==> logic to authentication
func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID := c.Param("id")
		bearer := c.Request.Header.Get("Authorization")

		AuthorID := routes.AllBlogs[ID].AuthorID

		if bearer == AuthorID {
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
