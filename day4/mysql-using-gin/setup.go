package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func SetUpDB() {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&Book{})
	DB = db
}
