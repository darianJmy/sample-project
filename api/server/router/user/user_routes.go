package user

import (
	"github.com/gin-gonic/gin"

	"sample-project/api/server/httputils"
	"sample-project/pkg/types"
)

type IdMeta struct {
	UserId int64 `uri:"userId" binding:"required"`
}

func (u *userRouter) createUser(c *gin.Context) {
	r := httputils.NewResponse()

	var (
		user types.User
		err  error
	)

	if err = c.ShouldBindJSON(&user); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}
	if r.Result, err = u.c.User().Create(c, &user); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (u *userRouter) updateUser(c *gin.Context) {
	r := httputils.NewResponse()

	var (
		idMeta IdMeta
		user   types.User
		err    error
	)

	if err = c.ShouldBindUri(&idMeta); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	if err = c.ShouldBindJSON(&user); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}
	if r.Result, err = u.c.User().Update(c, idMeta.UserId, &user); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (u *userRouter) deleteUser(c *gin.Context) {
	r := httputils.NewResponse()

	var (
		idMeta IdMeta
		err    error
	)

	if err = c.ShouldBindUri(&idMeta); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	if r.Result, err = u.c.User().Delete(c, idMeta.UserId); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (u *userRouter) getUser(c *gin.Context) {
	r := httputils.NewResponse()

	var (
		idMeta IdMeta
		err    error
	)

	if err = c.ShouldBindUri(&idMeta); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	if r.Result, err = u.c.User().Get(c, idMeta.UserId); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (u *userRouter) listUsers(c *gin.Context) {
	r := httputils.NewResponse()

	var err error

	if r.Result, err = u.c.User().List(c); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (u *userRouter) login(c *gin.Context) {
	r := httputils.NewResponse()

	var (
		user types.User
		err  error
	)

	if err = c.ShouldBindJSON(&user); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	if r.Result, err = u.c.User().Login(c, &user); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (u *userRouter) changePassword(c *gin.Context) {
	r := httputils.NewResponse()

	var (
		idMeta   IdMeta
		password types.Password
		err      error
	)

	if err = c.ShouldBindUri(&idMeta); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	if err = c.ShouldBindJSON(&password); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}
	if r.Result, err = u.c.User().ChangePassword(c, idMeta.UserId, &password); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (u *userRouter) resetPassword(c *gin.Context) {
	r := httputils.NewResponse()

	var (
		idMeta IdMeta
		err    error
	)

	if err = c.ShouldBindUri(&idMeta); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	if r.Result, err = u.c.User().ResetPassword(c, idMeta.UserId); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}
