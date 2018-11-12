package service

import (
	model "../model"
	orm "../db"
    "strings"
    // "strconv"
    "time"
    "fmt"
)

// 创建一个DemoOrder
func CreateDemoOrder(amount float64, orderid string, username string, status string, fileurl string) *model.DemoOrder {
    var do *model.DemoOrder = &model.DemoOrder{Amount:amount, Order_id:orderid, User_name:username, Status:status, File_url:fileurl}
    return do
}

// 添加
func Insert(do *model.DemoOrder) bool {
    orm.Eloquent.Create(do)
    return orm.Eloquent.NewRecord(*do)
}

// 修改
func Change(do *model.DemoOrder) bool {
	var target *model.DemoOrder = new(model.DemoOrder)
    orm.Eloquent.Where("Order_id = ? AND User_name = ?", (*do).Order_id, (*do).User_name).First(&target)
    if target == nil {
        fmt.Println("No such record to update")
        return true
    }
    orm.Eloquent.Model(&target).Updates(map[string]interface{}{"Amount": (*do).Amount, "Status": (*do).Status, "File_url": (*do).File_url})
    return false
}

// 删除
func Delete(order_id string, user_name string) bool {
	var target *model.DemoOrder = new(model.DemoOrder)
    orm.Eloquent.Where("Order_id = ? AND User_name = ?", order_id, user_name).First(&target)
    if target == nil {
        fmt.Println("No such record to delete")
        return true
    }
    orm.Eloquent.Where("Order_id = ? AND User_name = ?", order_id, user_name).Delete(model.DemoOrder{})
    return false
}

func SelectByCreatedAt(start time.Time, end time.Time) ([]model.DemoOrder){
    var result []model.DemoOrder
    query := orm.Eloquent.Where("created_at BETWEEN ? AND ?", start, end).Find(&result)
    if query.Error!=nil {
        fmt.Println("query failed")
        return nil
    }
    fmt.Println(result)
    return result
}


// 模糊查找，可以限定按照金额或者创建时间排序
func FuzzySearch(keyword string, sortby string, desc bool) (allrecord []model.DemoOrder, length int) {
    if !strings.EqualFold(sortby, "created_at") && !strings.EqualFold(sortby, "amount") {
        fmt.Println("Order condition is invalid.")
        return nil,1
    }
    if desc {
        sortby = sortby + " DESC"
    }
    keyword = "%"+keyword+"%"
    query := orm.Eloquent.Where("User_name LIKE ?", keyword).Or("Status LIKE ?", keyword).Or("File_url LIKE ?", keyword).Or("Order_id LIKE ?", keyword).Order(sortby, true).Find(&allrecord)
    if query.Error!=nil {
        fmt.Println("query failed")
        return nil, 0
    }
    length = len(allrecord)
    return allrecord, length
}

