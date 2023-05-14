package database

import (
	"fmt"
	"go_todo/pkg/entities"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	MYSQL_HOST := os.Getenv("MYSQL_HOST")
	MYSQL_USER := os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	MYSQL_DBNAME := os.Getenv("MYSQL_DBNAME")
	dsn := fmt.Sprintf("%v:%v@tcp(%v:3306)/%v?charset=utf8mb4&parseTime=True&loc=Local", MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_DBNAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func MigrateDDL(db *gorm.DB) {
	db.AutoMigrate(&entities.Activity{})
}
