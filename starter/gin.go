// this package contains all starter, e.g. go-micro, gin...
package starter

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"time"
	"github.com/doublesouth/gin-scaffold/middleware"
	"github.com/doublesouth/gin-scaffold/conf"
	"github.com/doublesouth/gin-scaffold/router"
)

func Gin() {
	gin.SetMode(conf.Conf.Application.GinMode)
	r := gin.New()
	r.Use(
		middleware.RequestId(),
		middleware.RequestAndResponseLogger(1*time.Second),
		//middleware.ErrorHandler(),
		gin.Recovery(),
	)
	router.V1(r)

	glog.Info("start gin...")
	r.Run(conf.Conf.Server.Bind)
}
