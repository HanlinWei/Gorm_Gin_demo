package db

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
	// model "../model"
)

var Eloquent *gorm.DB

func init() {
    var err error
	Eloquent, err := gorm.Open("sqlite3", "./db/test.db")
	fmt.Println("Open database ", Eloquent)
	if err != nil {
		panic("failed to connect database")
	}
	defer Eloquent.Close()

	// Eloquent.AutoMigrate(&model.User{})
}