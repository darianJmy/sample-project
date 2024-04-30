package info

import (
	"github.com/gin-gonic/gin"
	"sample-project/cmd/app/options"
	"sample-project/pkg/controller"
)

type infoRouter struct {
	c controller.SampleInterface
}

func NewRouter(o *options.ServerRunOptions) {
	router := &infoRouter{
		o.Control,
	}

	router.initRoutes(o.HttpEngine)
}

func (ir *infoRouter) initRoutes(httpEngine *gin.Engine) {
	infoRoute := httpEngine.Group("/info")
	{
		infoRoute.POST("", ir.createInfo)
		infoRoute.PUT("/:infoId", ir.updateInfo)
		infoRoute.DELETE("/:infoId", ir.deleteInfo)
		infoRoute.GET("/:infoId", ir.getInfo)
		infoRoute.GET("", ir.listInfos)
	}
}
