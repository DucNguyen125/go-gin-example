package models

import (
	"gorm.io/gorm"
)

func AutoMigrateTable(db *gorm.DB) error {
	var err error
	if err = db.AutoMigrate(&User{}); err != nil {
		return err
	}
	if err = db.AutoMigrate(&Order{}); err != nil {
		return err
	}
	err = db.AutoMigrate(&Product{})
	return err
}
