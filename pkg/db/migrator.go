package db

import (
	"gorm.io/gorm"
)

type migrator struct {
	db *gorm.DB
}

// AutoMigrate 自动创建指定模型的数据库表结构
func (m *migrator) AutoMigrate() error {
	dst := []interface{}{}

	return m.CreateTables(dst...)
}

func (m *migrator) CreateTables(dst ...interface{}) error {
	for _, d := range dst {
		if m.db.Migrator().HasTable(d) {
			continue
		}
		if err := m.db.Migrator().CreateTable(d); err != nil {
			return err
		}
	}

	return nil
}

func newMigrator(db *gorm.DB) *migrator {
	return &migrator{db}
}
