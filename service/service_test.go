package service

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    _ "../db"
	model "../model"
    "testing"
    "time"
)

var sample model.DemoOrder = model.DemoOrder{gorm.Model{}, 11.22, "test-Order_id", "test-User_name", "test-Status", "test-File_url"}

func Test_CreateDemoOrder(t *testing.T) {
    created := CreateDemoOrder(11.22, "Order_id", "User_name", "Status", "File_url")
    if (*created) != sample {
        t.Error("Test_CreateDemoOrder 测试失败")
    }
}

func Test_Insert(t *testing.T) {
	created := Insert(&sample)
    if created {
		t.Error("Test_Insert 测试失败\n")
	}
}

func Test_Change(t *testing.T) {
	created := Change(&sample)
    if created {
		t.Error("Test_Change 测试失败\n")
	}
}

func Test_Delete(t *testing.T) {
	created := Delete(sample.Order_id, sample.User_name)
    if created {
		t.Error("Test_Delete 测试失败\n")
	}
}

func Test_SelectByCreatedAt(t *testing.T) {
    x := "2017-02-27 17:30:20"
    p, _ := time.Parse(x, x)
    result := SelectByCreatedAt(p, time.Now())
    if result == nil {
        t.Error("Test_SelectByCreatedAt 测试失败\n")
    }
}