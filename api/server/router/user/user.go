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

		// login
		userRoute.POST("/login", u.login)
		// change password

		userRoute.PUT("/change/password/:userId", u.changePassword)
		userRoute.PUT("/reset/password/:userId", u.resetPassword)

		userRoute.POST("/:userId/roles", u.userBindRoles)
		userRoute.DELETE("/:userId/roles", u.userUnBindRoles)

	}
}
