package db

import (
	"gorm.io/gorm"
)

type ShareDaoFactory interface {
}

type shareDaoFactory struct {
	db *gorm.DB
}

func (f *shareDaoFactory) User() UserInterface { return newUser(f.db) }

func NewDaoFactory(db *gorm.DB, migrate bool) (ShareDaoFactory, error) {
	if migrate {
		// automatically create the database table structure of the specified model
		if err := newMigrator(db).AutoMigrate(); err != nil {
			return nil, err
		}
	}

	return &shareDaoFactory{
		db: db,
	}, nil
}
