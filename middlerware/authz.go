package middlerware

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func NewAuthorizer(e *casbin.Enforcer) gin.HandlerFunc {
	a := &BasicAuthorizer{enforcer: e}

	return func(c *gin.Context) {
		if !a.CheckPermission(c.Request) {
			a.RequirePermission(c)
		}
	}
}

type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

func (a *BasicAuthorizer) CheckPermission(r *http.Request) bool {
	uid := r.Header.Get("x-tif-uid")
	method := r.Method
	path := r.URL.Path

	allowed, err := a.enforcer.Enforce(uid, path, method)
	if err != nil {
		panic(err)
	}

	return allowed
}

func (a *BasicAuthorizer) RequirePermission(c *gin.Context) {
	c.AbortWithStatus(403)
	c.JSON(403, gin.H{
		"message": "authz fail",
	})
}
