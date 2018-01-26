package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"github.com/doublesouth/gin-scaffold/common"
)

// middleware for generate reqeust id
func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get("X-Request-Id")
		if requestId == "" {
			requestId = uuid.NewV4().String()
		}

		// add requestId to context, next logic can use this
		c.Set(common.ContextRequestId, requestId)
		// add requestId to response header
		c.Writer.Header().Set("X-Request-Id", requestId)

		c.Next()
	}
}
