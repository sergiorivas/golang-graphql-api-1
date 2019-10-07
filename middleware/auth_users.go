package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
)

func AddUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "PTOUser", "admin-test-user")
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
