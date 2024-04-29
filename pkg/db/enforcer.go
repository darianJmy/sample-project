package db

import (
	"context"
	"strconv"

	"github.com/casbin/casbin/v2"
	csmodel "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"

	"sample-project/pkg/db/model"
)

type csEnforcer struct {
	enforcer *casbin.Enforcer
	db       *gorm.DB
}

// AddRoleForUser 分配用户角色
func (c *csEnforcer) AddRoleForUser(ctx context.Context, userid int64, roleIds []int64) error {
	uidStr := strconv.FormatInt(userid, 10)

	ok, err := c.enforcer.DeleteRolesForUser(uidStr)
	if err != nil || !ok {
		return err
	}

	for _, v := range roleIds {
		ok, err = c.enforcer.AddRoleForUser(uidStr, strconv.FormatInt(v, 10))
		if err != nil || !ok {
			return err
		}
	}

	return nil
}

// SetRolePermission 设置角色权限
func (c *csEnforcer) SetRolePermission(ctx context.Context, roleId int64, menus *[]model.Menu) (bool, error) {
	ok, err := c.enforcer.DeletePermissionsForUser(strconv.FormatInt(roleId, 10))
	if err != nil || !ok {
		return false, err
	}

	_, err = c.setRolePermission(roleId, menus)
	if err != nil {
		return false, err
	}
	return true, nil
}

// 设置角色权限
func (c *csEnforcer) setRolePermission(roleId int64, menus *[]model.Menu) (bool, error) {
	for _, v := range *menus {
		if v.MenuType == 2 || v.MenuType == 3 {
			ok, err := c.enforcer.AddPermissionForUser(strconv.FormatInt(roleId, 10), v.URL, v.Method)
			if err != nil || !ok {
				return ok, err
			}
		}
	}

	return false, nil
}

// DeleteRole 删除角色
func (c *csEnforcer) DeleteRole(ctx context.Context, roleId int64) error {
	ok, err := c.enforcer.DeletePermissionsForUser(strconv.FormatInt(roleId, 10))
	if err != nil || !ok {
		return err
	}

	ok, err = c.enforcer.DeleteRole(strconv.FormatInt(roleId, 10))
	if err != nil || !ok {
		return err
	}

	return nil
}

// DeleteRolePermission 删除角色权限
func (c *csEnforcer) DeleteRolePermission(ctx context.Context, resource ...string) error {
	ok, err := c.enforcer.DeletePermission(resource...)
	if err != nil || !ok {
		return err
	}

	return nil
}

func newEnforcer(db *gorm.DB) (*csEnforcer, error) {
	rbacRules :=
		`
	[request_definition]
	r = sub, obj, act
	
	[policy_definition]
	p = sub, obj, act
	
	[role_definition]
	g = _, _

	[policy_effect]
	e = some(where (p.eft == allow))
	
	[matchers]
	m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
	`
	// 加载鉴权规则
	m, err := csmodel.NewModelFromString(rbacRules)
	if err != nil {
		return nil, err
	}

	// 建表
	adapter, err := gormadapter.NewAdapterByDBWithCustomTable(db, &model.Rule{}, "rules")
	if err != nil {
		return nil, err
	}

	enforcer, err := casbin.NewEnforcer(m, adapter)
	// 加载权限
	if err = enforcer.LoadPolicy(); err != nil {
		return nil, err
	}

	return &csEnforcer{
		db:       db,
		enforcer: enforcer,
	}, nil
}
