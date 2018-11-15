package service

import (
	"fmt"
	"log"
	http "net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	orm "../db"
	model "../model"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

// 创建一个DemoOrder
func CreateDemoOrder(amount float64, orderid string, username string, status string, fileurl string) *model.DemoOrder {
	var do *model.DemoOrder = &model.DemoOrder{Amount: amount, Order_id: orderid, User_name: username, Status: status, File_url: fileurl}
	return do
}

// 添加
func Insert(do *model.DemoOrder) error {
	return orm.Eloquent.Create(do).Error
}

// 修改
func Change(do *model.DemoOrder) error {
	var target *model.DemoOrder = new(model.DemoOrder)
	err := orm.Eloquent.Where("Order_id = ? AND User_name = ?", (*do).Order_id, (*do).User_name).First(&target).Error
	if target == nil || err != nil {
		log.Println("No such record to update: ", err.Error())
		return err
	}
	if (*do).Amount > 0 {
		err = orm.Eloquent.Model(&target).Update("Amount", (*do).Amount).Error
		if err != nil {
			log.Println("Amount update failed: ", err.Error())
			return err
		}
	}
	if len((*do).Status) > 0 {
		err = orm.Eloquent.Model(&target).Update("Status", (*do).Status).Error
		if err != nil {
			log.Println("Status update failed: ", err.Error())
			return err
		}
	}
	if len((*do).File_url) > 0 {
		err = orm.Eloquent.Model(&target).Update("File_url", (*do).File_url).Error
		if err != nil {
			log.Println("File_url update failed: ", err.Error())
			return err
		}
	}
	return nil
}

// 删除
func Delete(order_id string, user_name string) error {
	var target *model.DemoOrder = new(model.DemoOrder)
	err := orm.Eloquent.Where("Order_id = ? AND User_name = ?", order_id, user_name).First(&target).Error
	if target == nil || err != nil {
		log.Println("No such record to delete: ", err)
		return err
	}
	return orm.Eloquent.Where("Order_id = ? AND User_name = ?", order_id, user_name).Delete(model.DemoOrder{}).Error
}

func SelectByCreatedAt(start time.Time, end time.Time) []model.DemoOrder {
	var result []model.DemoOrder
	query := orm.Eloquent.Where("created_at BETWEEN ? AND ?", start, end).Find(&result)
	if query.Error != nil {
		log.Println("query failed")
		return nil
	}
	log.Println(result)
	return result
}

// 模糊查找，可以限定按照金额或者创建时间排序
func FuzzySearch(keyword string, sortby string, desc bool) (allrecord []model.DemoOrder, length int, err error) {
	if !strings.EqualFold(sortby, "created_at") && !strings.EqualFold(sortby, "amount") {
		log.Println("Order condition is invalid.")
		return nil, 1, nil
	}
	if desc {
		sortby = sortby + " DESC"
	}
	keyword = "%" + keyword + "%"
	query := orm.Eloquent.Where("User_name LIKE ?", keyword).Or("Status LIKE ?", keyword).Or("File_url LIKE ?", keyword).Or("Order_id LIKE ?", keyword).Order(sortby, true).Find(&allrecord)
	if query.Error != nil {
		log.Println("query failed")
		return nil, 0, query.Error
	}
	length = len(allrecord)
	return allrecord, length, nil
}

// 数据以excel形式导出来
func ExportExcel() string {

	// 获取数据库中所有数据
	x := "2017-02-27 17:30:20"
	p, _ := time.Parse(x, x)
	result := SelectByCreatedAt(p, time.Now())

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	// 创建excel表格
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		log.Printf(err.Error())
	}
	// 写入数据
	for _, record := range result {
		row = sheet.AddRow()
		t := reflect.TypeOf(record)
		v := reflect.ValueOf(record)
		for k := 0; k < t.NumField(); k++ {
			cell = row.AddCell()
			cell.Value = fmt.Sprintf("%v", v.Field(k).Interface())
			// cell.value = log.Sprintf("%s -- %v \n", t.Field(k).Name, v.Field(k).Interface())
		}
	}
	// 关闭文件
	excel := "./data.xlsx"
	err = file.Save("data.xlsx")
	if err != nil {
		log.Printf(err.Error())
	}
	return excel
}

// 解析url参数
func SolveParam(c *gin.Context) (amount float64, order_id string, user_name string, status string, file_url string, err error) {
	amount, err = strconv.ParseFloat(c.DefaultQuery("amount", "-1"), 64)
	order_id = c.Query("order_id")
	user_name = c.Query("user_name")
	status = c.Query("status")
	file_url = c.Query("file_url")
	return amount, order_id, user_name, status, file_url, err
}

// 添加和修改数据的基本操作，只是替换其中的service接口
func BasicOperation(c *gin.Context, service_func func(do *model.DemoOrder) error, task string) {
	amount, order_id, user_name, status, file_url, err := SolveParam(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "amount formate false" + err.Error(),
		})
		return
	}

	new_record := CreateDemoOrder(amount, order_id, user_name, status, file_url)

	failed := service_func(new_record)
	if failed != nil {
		log.Println(failed)
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": task + "失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":      1,
		"message":   task + "成功",
		"order_id":  order_id,
		"user_name": user_name,
		"status":    status,
		"file_url":  file_url,
		"amount":    amount,
	})
}
