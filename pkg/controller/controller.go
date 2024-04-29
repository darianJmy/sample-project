package controller

import (
	"sample-project/cmd/app/config"
	"sample-project/pkg/controller/menu"
	"sample-project/pkg/controller/role"
	"sample-project/pkg/controller/user"
	"sample-project/pkg/db"
)

type SampleInterface interface {
	user.UserGetter
	role.RoleGetter
	menu.MenuGetter
}

type sample struct {
	cc      config.Config
	factory *db.ShareDaoFactory
}

func (s *sample) User() user.Interface { return user.NewUser(s.cc, s.factory) }

func (s *sample) Role() role.Interface { return role.NewRole(s.factory) }

func (s *sample) Menu() menu.Interface { return menu.NewMenu(s.factory) }

func New(cfg config.Config, f *db.ShareDaoFactory) SampleInterface {
	return &sample{
		cc:      cfg,
		factory: f,
	}
}
