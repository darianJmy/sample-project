package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"

	"sample-project/cmd/app/options"
	"sample-project/pkg/controller"
)

type middleware struct {
	c        controller.SampleInterface
	enforcer *casbin.Enforcer
}

func NewMiddlewares(o *options.ServerRunOptions) {
	m := &middleware{
		c:        o.Control,
		enforcer: o.Factory.Enforcer.GetEnforcer(),
	}

	m.initMiddlewares(o.HttpEngine)
}

func (m *middleware) initMiddlewares(httpEngine *gin.Engine) {
	httpEngine.Use(
		m.cors(),
		m.authentication(),
	)
}
