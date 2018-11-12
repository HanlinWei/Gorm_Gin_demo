package handler

import (
    "github.com/gin-gonic/gin"
	model "../model"
	service "../service"
    http "net/http"
    "strconv"
    "strings"
    "time"
    // "fmt"
)

// 列出所有数据
func ListAll(c *gin.Context) {
    x := "2017-02-27 17:30:20"
    p, _ := time.Parse(x, x)
    result := service.SelectByCreatedAt(p, time.Now())
    c.JSON(http.StatusOK, gin.H{
        "code": 1,
        "message": "列出所有数据",
        "data": result,
    })
}

// 模糊查找
func FuzzySearch(c *gin.Context) {
    keyword := strings.Split(c.Param("keyword"), ":keyword")[0]
    sortby := strings.Split(c.Param("sortby"), ":sortby")[0]
    desc, err := strconv.ParseBool(strings.Split(c.Param("desc"), ":desc")[0])
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "code": 1,
            "message": "参数desc不合法，应为true或false",
        })
        return
    }
    result, length := service.FuzzySearch(keyword, sortby, desc)
    if result == nil && length == 0{
        c.JSON(http.StatusOK, gin.H{
            "code": 1,
            "message": "查询失败",
        })
        return
    }
    if result == nil && length == 1{
        c.JSON(http.StatusOK, gin.H{
            "code": 1,
            "message": "参数sortby不合法，应为createdat或amount",
        })
        return
    }
    if length == 0 {
        c.JSON(http.StatusOK, gin.H{
            "code": 1,
            "message": "未找到匹配",
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "code": 1,
        "message": "查询成功",
        "record number": length,
        "data": result,
    })
}

// 解析url参数
func solveParam(c *gin.Context) (amount float64, order_id string, user_name string, status string, file_url string, err error) {
    amount, err = strconv.ParseFloat(strings.Split(c.Param("amount"), ":amount")[0], 64)
    order_id = strings.Split(c.Param("order_id"), ":order_id")[0]
    user_name = strings.Split(c.Param("user_name"), ":user_name")[0]
    status = strings.Split(c.Param("status"), ":status")[0]
    file_url = strings.Split(c.Param("file_url"), ":file_url")[0]
    return amount, order_id, user_name, status, file_url, err
}

// 添加和修改数据的基本操作，只是替换其中的service接口
func basicOperation(c *gin.Context,service_func func(do *model.DemoOrder) (bool), task string) {
    amount, order_id, user_name, status, file_url, err:= solveParam(c)
    if err != nil {
		c.JSON(http.StatusOK, gin.H{
            "code":  0,
            "message": "amount fserviceat false" + err.Error(),
        })
        return
    }

    new_record := service.CreateDemoOrder(amount, order_id, user_name, status, file_url)

    failed := service_func(new_record)
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
    basicOperation(c, service.Insert, "创建")
}

// 修改数据
func Update(c *gin.Context) {
    basicOperation(c, service.Change, "更新")
}

// 删除数据
func Delete(c *gin.Context) {
    order_id := strings.Split(c.Param("order_id"), ":order_id")[0]
    user_name := strings.Split(c.Param("user_name"), ":user_name")[0]

    failed := service.Delete(order_id, user_name)
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