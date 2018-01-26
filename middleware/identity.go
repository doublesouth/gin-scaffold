package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/doublesouth/gin-scaffold/common"
)

// middleware for resolve the identity of this request
func Identity() gin.HandlerFunc {
	return func(c *gin.Context) {
		identity := new(AppIdentity)
		if _, err := identity.Resolve(c); err != nil {
			c.JSON(common.ErrorResult(err))
			c.Abort()
			return
		}

		// add userid to context, next logic can use this
		c.Set(common.ContextIdentity, identity)

		c.Next()
	}
}

type AppIdentity struct {
}

type IdentityResolver interface {
	Resolve(c *gin.Context) (IdentityResolver, error)
	GetIdentity() string
	GetIdentityInfo() interface{}
}

func (i *AppIdentity) Resolve(c *gin.Context) (IdentityResolver, error) {
	// TODO need to implement
	return nil, nil
}

func (i *AppIdentity) GetIdentity() string {
	// TODO need to implement
	return "mock-identity"
}

func (i *AppIdentity) GetIdentityInfo() interface{} {
	// TODO need to implement
	return nil
}
