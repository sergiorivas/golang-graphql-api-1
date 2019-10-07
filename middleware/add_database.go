package middleware

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

func AddDatabase(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "DBConnection", db)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GetConnection(ctx context.Context) *gorm.DB {
	return ctx.Value("DBConnection").(*gorm.DB)
}
