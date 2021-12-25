package db

import (
	"acmicpc_checker_v2_backend/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
	// Database init
	db, err = gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Student{})
	db.AutoMigrate(&model.Assignment{})
	db.AutoMigrate(&model.SolvedInfo{})
	db.AutoMigrate(&model.ClassInfo{})
	db.AutoMigrate(&model.ClassStudent{})
}

func Dbconnect() *gorm.DB {
	return db
}

func TestDbconnect() *gorm.DB {
	db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
