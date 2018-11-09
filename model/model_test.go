package model

import (
    "testing"
    "github.com/jinzhu/gorm"
)

var sample DemoOrder = DemoOrder{gorm.Model{}, 11.22, "Order_id", "User_name", "Status", "File_url"}

func TestCreateDemoOrder(t *testing.T) {
    created := CreateDemoOrder(11.22, "Order_id", "User_name", "Status", "File_url")
    if (*created) != sample {
        t.Error("TestCreateDemoOrder 测试失败")
    }
}

