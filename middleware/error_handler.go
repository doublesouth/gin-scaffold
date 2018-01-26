package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

// 处理context中的error
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()

		err := c.Errors.Last()
		if err != nil {
			// 打印错误堆栈信息
			glog.Errorf("err: %+v", err.Err)
		}
	}
}
