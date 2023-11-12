package middleware

import (
	"fisco/config"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func NewAuthorizer(e *casbin.Enforcer) gin.HandlerFunc {
	a := &BasicAuthorizer{enforcer: e}

	return func(c *gin.Context) {
		if !a.CheckPermission(c) {
			a.RequirePermission(c)
		}
	}
}

type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

func (a *BasicAuthorizer) GetRoleName(c *gin.Context) interface{} {
	session, err := config.Store.Get(c.Request, "sessionID")
	if err != nil {
		//config.Logger.Error(err.Error())
	}
	role := session.Values["role"]
	if role != nil && role != "" {
		return role
	}
	return "common"
}

func (a *BasicAuthorizer) CheckPermission(c *gin.Context) bool {
	role := a.GetRoleName(c)
	method := c.Request.Method
	path := c.Request.URL.Path
	enforce, err := a.enforcer.Enforce(role, path, method)
	if err != nil {
		//config.Logger.Error(err.Error())
	}
	return enforce
}

func (a *BasicAuthorizer) RequirePermission(c *gin.Context) {
	c.AbortWithStatus(403)
}
