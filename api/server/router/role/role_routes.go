package role

import (
	"github.com/gin-gonic/gin"

	"sample-project/api/server/httputils"
	"sample-project/pkg/types"
)

type IdMeta struct {
	RoleId int64 `uri:"roleId" binding:"required"`
}

func (ro *roleRouter) createRole(c *gin.Context) {
	r := httputils.NewResponse()

	var (
		role types.Role
		err  error
	)

	if err = c.ShouldBindJSON(&role); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	if r.Result, err = ro.c.Role().Create(c, &role); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (ro *roleRouter) updateRole(c *gin.Context) {
	r := httputils.NewResponse()

	var (
		idMeta IdMeta
		role   types.Role
		err    error
	)

	if err = c.ShouldBindUri(&idMeta); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	if err = c.ShouldBindJSON(&role); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}
	if r.Result, err = ro.c.Role().Update(c, idMeta.RoleId, &role); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (ro *roleRouter) deleteRole(c *gin.Context) {
	r := httputils.NewResponse()

	var (
		idMeta IdMeta
		err    error
	)

	if err = c.ShouldBindUri(&idMeta); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	if r.Result, err = ro.c.Role().Delete(c, idMeta.RoleId); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (ro *roleRouter) getRole(c *gin.Context) {
	r := httputils.NewResponse()

	var (
		idMeta IdMeta
		err    error
	)

	if err = c.ShouldBindUri(&idMeta); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	if r.Result, err = ro.c.Role().Get(c, idMeta.RoleId); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (ro *roleRouter) listRoles(c *gin.Context) {
	r := httputils.NewResponse()

	var err error

	if r.Result, err = ro.c.Role().List(c); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (ro *roleRouter) roleBindMenus(c *gin.Context) {
	r := httputils.NewResponse()

	var (
		idMeta IdMeta
		menus  types.Menus
		err    error
	)

	if err = c.ShouldBindUri(&idMeta); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	if err = c.ShouldBindJSON(&menus); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	if err = ro.c.Role().RoleBindMenu(c, idMeta.RoleId, menus.MenuIds); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (ro *roleRouter) roleUnBindMenus(c *gin.Context) {
	r := httputils.NewResponse()

	httputils.SetSuccess(c, r)
}
