package model

import "gorm.io/gorm"

// User 用户
type User struct {
	gorm.Model

	Name        string `gorm:"index:idx_name,unique" json:"name"`
	Password    string `gorm:"type:varchar(256)" json:"-"`
	Status      int8   `gorm:"type:tinyint" json:"status"`
	Role        int    `gorm:"type:tinyint" json:"role"`
	Email       string `gorm:"type:varchar(128)" json:"email"`
	Description string `gorm:"type:text" json:"description"`
	Extension   string `gorm:"type:text" json:"extension,omitempty"`
}

func (user *User) TableName() string { return "users" }

// Role 角色
type Role struct {
	gorm.Model

	Memo     string `gorm:"column:memo;size:128;" json:"memo" form:"memo"`                                     // 备注
	Name     string `gorm:"column:name;size:128;not null;unique_index:uk_role_name;;" json:"name" form:"name"` // 名称
	Sequence int    `gorm:"column:sequence;not null;" json:"sequence" form:"sequence"`                         // 排序值
	ParentID int64  `gorm:"column:parent_id;not null;" json:"parent_id" form:"parent_id"`                      // 父级ID
	Status   int8   `gorm:"column:status" json:"status" form:"status"`                                         // 0 表示禁用，1 表示启用
	Children []Role `gorm:"-"`
}

func (r *Role) TableName() string { return "roles" }

// Menu 菜单
type Menu struct {
	gorm.Model

	Status   int8   `gorm:"column:status;type:tinyint(1);not null;" json:"status" form:"status"`          // 状态(1:启用 2:不启用)
	Memo     string `gorm:"column:memo;size:128;" json:"memo,omitempty" form:"memo"`                      // 备注
	ParentID int64  `gorm:"column:parent_id;not null;" json:"parent_id,omitempty" form:"parent_id"`       // 父级ID
	URL      string `gorm:"column:url;size:128;" json:"url,omitempty" form:"url"`                         // 菜单URL
	Name     string `gorm:"column:name;size:128;not null;" json:"name" form:"name"`                       // 菜单名称
	Sequence int    `gorm:"column:sequence;not null;" json:"sequence" form:"sequence"`                    // 排序值
	MenuType int8   `gorm:"column:menu_type;type:tinyint(1);not null;" json:"menu_type" form:"menu_type"` // 菜单类型 1 左侧菜单,2 按钮, 3 非展示权限
	Icon     string `gorm:"column:icon;size:32;" json:"icon,omitempty" form:"icon"`                       // icon
	Method   string `gorm:"column:method;size:32;not null;" json:"method,omitempty" form:"method"`        // 操作类型 none/GET/POST/PUT/DELETE
	Children []Menu `gorm:"-" json:"children"`
}

func (m *Menu) TableName() string {
	return "menus"
}

// UserRole 用户绑定角色
type UserRole struct {
	gorm.Model

	UserID int64 `gorm:"column:user_id;unique_index:uk_user_role_user_id;not null;" json:"user_id"` // 管理员ID
	RoleID int64 `gorm:"column:role_id;unique_index:uk_user_role_user_id;not null;" json:"role_id"` // 角色ID
}

func (u *UserRole) TableName() string { return "user_roles" }

// RoleMenu 角色绑定菜单
type RoleMenu struct {
	gorm.Model

	RoleID int64 `gorm:"column:role_id;unique_index:uk_role_menu_role_id;not null;" json:"role_id"`  // 角色ID
	MenuID int64 `gorm:"column:menu_id;unique_index:uk_role_menu_role_id;not null;" json:"menu_id'"` // 菜单ID
}

func (m *RoleMenu) TableName() string {
	return "role_menus"
}

// Rule 规则，由 casbin 控制
type Rule struct {
	gorm.Model

	PType  string `json:"ptype" gorm:"column:ptype;size:100" description:"策略类型"`
	Role   string `json:"role" gorm:"column:v0;size:100" description:"角色"`
	Path   string `json:"path" gorm:"column:v1;size:100" description:"api路径"`
	Method string `json:"method" gorm:"column:v2;size:100" description:"访问方法"`
	V3     string `gorm:"column:v3;size:100"`
	V4     string `gorm:"column:v4;size:100"`
	V5     string `gorm:"column:v5;size:100"`
}

func (r *Rule) TableName() string { return "rules" }
