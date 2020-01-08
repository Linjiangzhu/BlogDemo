package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// MethodOverride override request method by indpect the hidden filed in form post
func MethodOverride(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "POST" {
			return
		}
		method := c.PostForm("_method")
		if method == "" {
			return
		}
		method = strings.ToLower(method)
		switch method {
		case "delete":
			c.Request.Method = "DELETE"
		case "put":
			c.Request.Method = "PUT"
		case "patch":
			c.Request.Method = "PATCH"
		default:
			return
		}
		r.HandleContext(c)
		return
	}
}
