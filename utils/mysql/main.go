package mysql

import (
	"fmt"
	"os"

	"example/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	connectString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_URI"),
		os.Getenv("MYSQL_DATABASE_NAME"))
	db, err := gorm.Open(mysql.Open(connectString), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}

func AutoMigrate() error {
	err := models.AutoMigrateTable(DB)
	return err
}
