package user

import (
	"github.com/gin-gonic/gin"
	"sample-project/cmd/app/options"
	"sample-project/pkg/controller"
)

type userRouter struct {
	c controller.SampleInterface
}

func NewRouter(o *options.ServerRunOptions) {
	router := &userRouter{
		o.Control,
	}

	router.initRoutes(o.HttpEngine)
}

func (u *userRouter) initRoutes(httpEngine *gin.Engine) {
	userRoute := httpEngine.Group("/users")
	{
		userRoute.POST("", u.createUser)
		userRoute.PUT("/:userId", u.updateUser)
		userRoute.DELETE("/:userId", u.deleteUser)
		userRoute.GET("/:userId", u.getUser)
		userRoute.GET("", u.listUsers)
	}
}
