package role

import (
	"github.com/gin-gonic/gin"

	"sample-project/cmd/app/options"
	"sample-project/pkg/controller"
)

type roleRouter struct {
	c controller.SampleInterface
}

func NewRouter(o *options.ServerRunOptions) {
	router := &roleRouter{
		o.Control,
	}

	router.initRoutes(o.HttpEngine)
}

func (ro *roleRouter) initRoutes(httpEngine *gin.Engine) {
	roleRoute := httpEngine.Group("/roles")
	{
		roleRoute.POST("", ro.createRole)
		roleRoute.PUT("/:roleId", ro.updateRole)
		roleRoute.DELETE("/:roleId", ro.deleteRole)
		roleRoute.GET("/:roleId", ro.getRole)
		roleRoute.GET("", ro.listRoles)

		roleRoute.POST("/:roleId/menus", ro.roleBindMenus)
		roleRoute.DELETE("/:roleId/menus", ro.roleUnBindMenus)
	}
}
