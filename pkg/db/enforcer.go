package db

import (
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

// GetEnforcer 获取直接操作
func (c *csEnforcer) GetEnforcer() *casbin.Enforcer {
	return c.enforcer
}

func (c *csEnforcer) UserBindRoles(userid int64, roleIds ...int64) error {
	uidStr := strconv.FormatInt(userid, 10)

	ok, err := c.enforcer.DeleteRolesForUser(uidStr)
	if err != nil || !ok {
		return err
	}

	for _, roleId := range roleIds {
		ok, err = c.enforcer.AddRoleForUser(uidStr, strconv.FormatInt(roleId, 10))
		if err != nil || !ok {
			return err
		}
	}

	return nil
}

func (c *csEnforcer) RoleBindMenus(roleId int64, menus []*model.Menu) (bool, error) {
	_, err := c.enforcer.DeletePermissionsForUser(strconv.FormatInt(roleId, 10))
	if err != nil {
		return false, err
	}

	for _, v := range menus {
		ok, err := c.enforcer.AddPermissionForUser(strconv.FormatInt(roleId, 10), v.URL, v.Method)
		if err != nil || !ok {
			return ok, err
		}
	}

	return true, nil
}

func (c *csEnforcer) DeleteRole(roleId int64) error {
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

func (c *csEnforcer) DeleteRolePermission(resource ...string) error {
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
