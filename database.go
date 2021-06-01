package main

import (
	"fmt"
	"github.com/Bastorx/test_lbc/logging"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func InitDatabase() (*gorm.DB, error) {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PWD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")

	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, database)
	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if err != nil {
		logging.Logger.Error(fmt.Sprintf("Cannot connect to database\n", err.Error()))
	}

	return db, err
}