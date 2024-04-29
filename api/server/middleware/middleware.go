package middleware

import (
	"github.com/gin-gonic/gin"

	"sample-project/cmd/app/options"
	"sample-project/pkg/controller"
)

type middleware struct {
	c controller.SampleInterface
}

func NewMiddlewares(o *options.ServerRunOptions) {
	m := &middleware{
		o.Control,
	}

	m.initMiddlewares(o.HttpEngine)
}

func (m *middleware) initMiddlewares(httpEngine *gin.Engine) {
	httpEngine.Use(
		m.cors(),
		m.authentication(),
	)
}
