package menu

import (
	"github.com/gin-gonic/gin"

	"sample-project/cmd/app/options"
	"sample-project/pkg/controller"
)

type menuRouter struct {
	c controller.SampleInterface
}

func NewRouter(o *options.ServerRunOptions) {
	router := &menuRouter{
		o.Control,
	}

	router.initRoutes(o.HttpEngine)
}

func (m *menuRouter) initRoutes(httpEngine *gin.Engine) {
	menuRoute := httpEngine.Group("/menus")
	{
		menuRoute.POST("", m.createMenu)
		menuRoute.PUT("/:menuId", m.updateMenu)
		menuRoute.DELETE("/:menuId", m.deleteMenu)
		menuRoute.GET("/:menuId", m.getMenu)
		menuRoute.GET("", m.listMenus)
	}
}
