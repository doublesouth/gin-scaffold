package router

import (
	"github.com/gin-gonic/gin"
	"github.com/doublesouth/gin-scaffold/middleware"
	"github.com/doublesouth/gin-scaffold/controller"
)

// v1版本路由
func V1(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.Use(middleware.Identity())
	{
		v1.GET("/ping", controller.Ping)
	}
}
