package user

import (
	"github.com/gin-gonic/gin"
	"sample-project/api/server/httputils"
)

func (u *userRouter) createUser(c *gin.Context) {
	r := httputils.NewResponse()

	httputils.SetSuccess(c, r)
}

func (u *userRouter) updateUser(c *gin.Context) {
	r := httputils.NewResponse()

	httputils.SetSuccess(c, r)
}

func (u *userRouter) deleteUser(c *gin.Context) {
	r := httputils.NewResponse()

	httputils.SetSuccess(c, r)
}

func (u *userRouter) getUser(c *gin.Context) {
	r := httputils.NewResponse()

	httputils.SetSuccess(c, r)
}

func (u *userRouter) listUsers(c *gin.Context) {
	r := httputils.NewResponse()

	httputils.SetSuccess(c, r)
}
