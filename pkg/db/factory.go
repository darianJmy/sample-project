package db

import (
	"gorm.io/gorm"

	"sample-project/pkg/db/model"
)

type ShareDaoFactory struct {
	User     *Curd[*model.User]
	Role     *Curd[*model.Role]
	UserRole *Curd[*model.UserRole]
	Menu     *Curd[*model.Menu]
	RoleMenu *Curd[*model.RoleMenu]

	Enforcer *csEnforcer
}

func NewDaoFactory(db *gorm.DB, migrate bool) (*ShareDaoFactory, error) {
	if migrate {
		// automatically create the database table structure of the specified model
		if err := newMigrator(db).CreateTables(
			&model.User{},
			&model.Role{},
			&model.UserRole{},
			&model.Menu{},
			&model.RoleMenu{},
		); err != nil {
			return nil, err
		}
	}

	enforcer, err := newEnforcer(db)
	if err != nil {
		return nil, err
	}

	return &ShareDaoFactory{
		User:     newCurd(&model.User{}, db),
		Role:     newCurd(&model.Role{}, db),
		UserRole: newCurd(&model.UserRole{}, db),
		Menu:     newCurd(&model.Menu{}, db),
		RoleMenu: newCurd(&model.RoleMenu{}, db),
		Enforcer: enforcer,
	}, nil
}
