package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open("mysql", "root:belikemee@tcp(127.0.0.1:3306)/try")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected!")

	db.AutoMigrate(&Book{})

	DB = db
}
