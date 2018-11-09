package service

import (
    "github.com/gin-gonic/gin"
	model "../model"
	orm "../db"
    http "net/http"
    "strconv"
    "strings"
    "time"
    // "fmt"
)

// 列出所有数据
func ListAll(c *gin.Context) {
    x := "2017-02-27 17:30:20"
    p, _ := time.Parse("2006-01-02 15:04:05", x)
    result := orm.SelectByCreatedAt(p, time.Now())
    c.JSON(http.StatusOK, gin.H{
        "code": 1,
        "message": "列出所有数据",
        "data": result,
    })
}

// 模糊查找
func FuzzySearch(c *gin.Context) {
    query := strings.Split(c.Param("query"), ":query")[0]
    
}


// 解析url参数
func solveParam(c *gin.Context) (amount float64, order_id string, user_name string, status string, file_url string, err error) {
    amount, err = strconv.ParseFloat(strings.Split(c.Param("amount"), ":amount")[0], 64)
    order_id = strings.Split(c.Param("order_id"), ":order_id")[0]
    user_name = strings.Split(c.Param("user_name"), ":user_name")[0]
    status = strings.Split(c.Param("status"), ":status")[0]
    file_url = strings.Split(c.Param("file_url"), ":file_url")[0]
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "code":  0,
            "message": "amount format false" + err.Error(),
        })
    }
    return amount, order_id, user_name, status, file_url, err
}

// 添加和修改数据的基本操作，只是替换其中的orm接口
func basicOperation(c *gin.Context,fun func(do *model.DemoOrder) (bool), task string) {
    amount, order_id, user_name, status, file_url, err:= solveParam(c)
    if err != nil {
        return
    }

    new_record := model.CreateDemoOrder(amount, order_id, user_name, status, file_url)

    failed := fun(new_record)
    if failed {
        c.JSON(http.StatusOK, gin.H{
            "code":  0,
            "message": task + "失败",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "code":  1,
        "message": task + "成功",
        "order_id": order_id,
        "user_name": user_name,
        "status": status,
        "file_url": file_url,
        "amount": amount,
    })
}

// 添加数据
func Store(c *gin.Context) {
    basicOperation(c, orm.Insert, "创建")
}

// 修改数据
func Update(c *gin.Context) {
    basicOperation(c, orm.Change, "更新")
}

// 删除数据
func Delete(c *gin.Context) {
    order_id := strings.Split(c.Param("order_id"), ":order_id")[0]
    user_name := strings.Split(c.Param("user_name"), ":user_name")[0]

    failed := orm.Delete(order_id, user_name)
    if failed {
        c.JSON(http.StatusOK, gin.H{
            "code":  0,
            "message": "删除失败",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "code":  1,
        "message": "删除成功",
        "order_id": order_id,
        "user_name": user_name,
    })
}