package menu

import (
	"github.com/gin-gonic/gin"

	"sample-project/api/server/httputils"
	"sample-project/pkg/types"
)

type IdMeta struct {
	MenuId int64 `uri:"menuId" binding:"required"`
}

func (m *menuRouter) createMenu(c *gin.Context) {
	r := httputils.NewResponse()

	var (
		menu types.Menu
		err  error
	)

	if err = c.ShouldBindJSON(&menu); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}
	if r.Result, err = m.c.Menu().Create(c, &menu); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (m *menuRouter) updateMenu(c *gin.Context) {
	r := httputils.NewResponse()

	var (
		idMeta IdMeta
		menu   types.Menu
		err    error
	)

	if err = c.ShouldBindUri(&idMeta); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	if err = c.ShouldBindJSON(&menu); err != nil {
		httputils.SetFailed(c, r, err)
		return

	}
	if r.Result, err = m.c.Menu().Update(c, idMeta.MenuId, &menu); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (m *menuRouter) deleteMenu(c *gin.Context) {
	r := httputils.NewResponse()

	var (
		idMeta IdMeta
		err    error
	)

	if err = c.ShouldBindUri(&idMeta); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	if r.Result, err = m.c.Menu().Delete(c, idMeta.MenuId); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (m *menuRouter) getMenu(c *gin.Context) {
	r := httputils.NewResponse()

	var (
		idMeta IdMeta
		err    error
	)

	if err = c.ShouldBindUri(&idMeta); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	if r.Result, err = m.c.Menu().Get(c, idMeta.MenuId); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (m *menuRouter) listMenus(c *gin.Context) {
	r := httputils.NewResponse()

	var err error

	if r.Result, err = m.c.Menu().List(c); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}
