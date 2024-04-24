package db

import (
	"context"
	"gorm.io/gorm"
	"sample-project/pkg/db/model"
)

type UserInterface interface {
	Create(ctx context.Context, object *model.User) (*model.User, error)
}

type user struct {
	db *gorm.DB
}

func (u *user) Create(ctx context.Context, object *model.User) (*model.User, error) {
	if err := u.db.Create(object).Error; err != nil {
		return nil, err
	}
	return object, nil
}

func newUser(db *gorm.DB) *user {
	return &user{db}
}
