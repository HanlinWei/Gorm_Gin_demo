package db

import (
	"log"

	model "../model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "time"
	// "strings"
)

var Eloquent *gorm.DB

func init() {
	var err error
	Eloquent, err = gorm.Open("mysql", "qydev:654321@/qiyuan_demo?charset=utf8&parseTime=True&loc=Local")
	log.Println("Try to open database.")
	if err != nil {
		defer Eloquent.Close()
		panic("failed to connect database: " + err.Error())
	}
	log.Println("Successfully open database.")
	Eloquent.AutoMigrate(&model.DemoOrder{})
}

func CloseDB() {
	log.Println("Closing database.")
	if err := Eloquent.Close(); err != nil {
		log.Println("Close database failed: ", err.Error())
	}
	log.Println("Close database successfully.")
	return
}
