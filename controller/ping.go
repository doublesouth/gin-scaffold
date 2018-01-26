package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/doublesouth/gin-scaffold/common"
	"github.com/doublesouth/gin-scaffold/middleware"
	"github.com/golang/glog"
)

func Ping(c *gin.Context) {
	identity, _ := c.Get(common.ContextIdentity)
	i := identity.(middleware.IdentityResolver)
	glog.Infof("identity: %s", i.GetIdentity())

	c.JSON(http.StatusOK, common.Success())
}
