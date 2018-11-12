package service

import (
	model "../model"
	orm "../db"
    "strings"
    // "strconv"
    "time"
    "fmt"
    "github.com/tealeg/xlsx"
    "reflect"
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
    if (*do).Amount > 0 {
        orm.Eloquent.Model(&target).Update("Amount", (*do).Amount)
    }
    if len((*do).Status) > 0 {
        orm.Eloquent.Model(&target).Update("Status", (*do).Status)
    }
    if len((*do).File_url) > 0 {
        orm.Eloquent.Model(&target).Update("File_url", (*do).File_url)
    }
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

// 数据以excel形式导出来
func ExportExcel() {

    // 获取数据库中所有数据
    x := "2017-02-27 17:30:20"
    p, _ := time.Parse(x, x)
    result := SelectByCreatedAt(p, time.Now())

    var file *xlsx.File
    var sheet *xlsx.Sheet
    var row *xlsx.Row
    var cell *xlsx.Cell
    var err error

    file = xlsx.NewFile()
    sheet, err = file.AddSheet("Sheet1")
    if err != nil {
        fmt.Printf(err.Error())
    }
    for _, record := range result {
        row = sheet.AddRow()
        t := reflect.TypeOf(record)
        v := reflect.ValueOf(record)
        for k := 0; k < t.NumField(); k++ {
            cell = row.AddCell()
            cell.Value = fmt.Sprintf("%v", v.Field(k).Interface())
			// cell.value = fmt.Sprintf("%s -- %v \n", t.Field(k).Name, v.Field(k).Interface())   
	    }
    }
    err = file.Save("MyXLSXFile.xlsx")
    if err != nil {
        fmt.Printf(err.Error())
    }
}

