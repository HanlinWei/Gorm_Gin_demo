package db

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
    model "../model"
    // "time"
    // "strings"
)

var Eloquent *gorm.DB

func init() {
    var err error
	Eloquent, err = gorm.Open("sqlite3", "./db/test.db")
	fmt.Println("Try to open database ", Eloquent)
	if err != nil {
		defer Eloquent.Close()
		panic("failed to connect database: " + err.Error())
	}
    fmt.Println("Successfully open database ", Eloquent)
	Eloquent.AutoMigrate(&model.DemoOrder{})
}

func CloseDB() {
    Eloquent.Close()
    return
}
