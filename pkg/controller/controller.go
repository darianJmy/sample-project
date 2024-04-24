package controller

import (
	"sample-project/cmd/app/config"
	"sample-project/pkg/controller/user"
	"sample-project/pkg/db"
)

type SampleInterface interface {
	user.UserGetter
}

type sample struct {
	cc      config.Config
	factory db.ShareDaoFactory
}

func (s *sample) User() user.Interface { return user.NewUser(s.cc, s.factory) }

func New(cfg config.Config, f db.ShareDaoFactory) SampleInterface {
	return &sample{
		cc:      cfg,
		factory: f,
	}
}
