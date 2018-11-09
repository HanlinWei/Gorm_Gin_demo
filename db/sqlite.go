package db

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
    model "../model"
    "time"
    "strings"
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
}

// 添加
func Insert(do *model.DemoOrder) bool {
    Eloquent.Create(do)
    return Eloquent.NewRecord(*do)
}

// 修改
func Change(do *model.DemoOrder) bool {
	var target *model.DemoOrder = new(model.DemoOrder)
    Eloquent.Where("Order_id = ? AND User_name = ?", (*do).Order_id, (*do).User_name).First(&target)
    if target == nil {
        fmt.Println("No such record to update")
        return true
    }
    (*target).Amount = (*do).Amount
    (*target).Status = (*do).Status
    (*target).File_url = (*do).File_url
    return false
}

// 删除
func Delete(order_id string, user_name string) bool {
	var target *model.DemoOrder = new(model.DemoOrder)
    Eloquent.Where("Order_id = ? AND User_name = ?", order_id, user_name).First(&target)
    if target == nil {
        fmt.Println("No such record to delete")
        return true
    }
    Eloquent.Where("Order_id = ? AND User_name = ?", order_id, user_name).Delete(model.DemoOrder{})
    return false
}

func SelectByCreatedAt(start time.Time, end time.Time) ([]model.DemoOrder){
    var result []model.DemoOrder
    query := Eloquent.Where("created_at BETWEEN ? AND ?", start, end).Find(&result)
    if query.Error!=nil {
        fmt.Println("query failed")
        return nil
    }
    fmt.Println(result)
    return result
}

func FuzzySearch(keyword string) (result *model.DemoOrder, length int) {
    length = 0
    var start *model.DemoOrder
    var allrecord []model.DemoOrder
    
    query := Eloquent.Where("ID >= ?", 0).Find(&allrecord)
    if query.Error!=nil {
        fmt.Println("query failed")
        return nil, 0
    }

    result = len(allrecord)

    var record string
    for i := range allrecord {
        record = fmt.Sprintf("%s", allrecord[i])
        if strings.Contains(record, keyword) {

        }
    }
}
