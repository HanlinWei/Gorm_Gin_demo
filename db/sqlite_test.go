package db

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
	model "../model"
	"testing"
)

var sample model.DemoOrder = model.DemoOrder{gorm.Model{}, 11.22, "Order_id", "User_name", "Status", "File_url"}

func TestInsert(t *testing.T) {
	err := Insert(sample)
    if err == nil {
		t.Error("TestInsert 测试失败\nid = "+ err.Error())
	}
	CloseDB()
}