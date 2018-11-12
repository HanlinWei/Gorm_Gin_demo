package service

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
	model "../model"
	"testing"
)

var sample model.DemoOrder = model.DemoOrder{gorm.Model{}, 11.22, "Order_id", "User_name", "Status", "File_url"}

func TestInsert(t *testing.T) {
	created := Insert(&sample)
    if created {
		t.Error("TestInsert 测试失败\n")
	}
}

func TestCreateDemoOrder(t *testing.T) {
    created := CreateDemoOrder(11.22, "Order_id", "User_name", "Status", "File_url")
    if (*created) != sample {
        t.Error("TestCreateDemoOrder 测试失败")
    }
}