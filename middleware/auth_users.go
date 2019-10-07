package middleware

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(msg string) interface{} {
	str := fmt.Sprintf("{\"errors\": [{ \"message\": \"%s\" }], \"data\": null }", msg)
	var res interface{}
	json.Unmarshal([]byte(str), &res)
	return res
}

func AddUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "PTOUser", "admin-test-user")
		c.Request = c.Request.WithContext(ctx)
		c.Next()

		// res := ErrorResponse("Bad Request")
		// c.JSON(http.StatusBadRequest, res)
		// c.Abort()
	}
}
