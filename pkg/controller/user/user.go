package user

import (
	"context"
	"sample-project/cmd/app/config"
	"sample-project/pkg/db"
	"sample-project/pkg/db/model"
	"sample-project/pkg/types"
)

type UserGetter interface {
	User() Interface
}

type Interface interface {
	Create(ctx context.Context, user *types.User) error
}

type user struct {
	cc      config.Config
	factory db.ShareDaoFactory
}

func (u *user) Create(ctx context.Context, user *types.User) error {
	if _, err := u.factory.User().Create(ctx, &model.User{
		Name:     user.Name,
		Password: user.Password,
	}); err != nil {
		return err
	}

	return nil
}

func NewUser(cfg config.Config, f db.ShareDaoFactory) *user {
	return &user{
		cc:      cfg,
		factory: f,
	}
}
