package handler

import (
    "github.com/gin-gonic/gin"
	model "../model"
	service "../service"
    http "net/http"
    "strconv"
    "log"
    "time"
    "os"
    "io"
    // "strings"
    "fmt"
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
    keyword := c.Query("keyword")
    sortby := c.Query("sortby")
    desc, err := strconv.ParseBool(c.Query("desc"))
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
    amount, err = strconv.ParseFloat(c.DefaultQuery("amount", "-1"), 64)
    order_id = c.Query("order_id")
    user_name = c.Query("user_name")
    status = c.Query("status")
    file_url = c.Query("file_url")
    return amount, order_id, user_name, status, file_url, err
}

// 添加和修改数据的基本操作，只是替换其中的service接口
func basicOperation(c *gin.Context,service_func func(do *model.DemoOrder) (bool), task string) {
    amount, order_id, user_name, status, file_url, err:= solveParam(c)
    if err != nil {
		c.JSON(http.StatusOK, gin.H{
            "code":  0,
            "message": "amount formate false" + err.Error(),
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
    order_id := c.Query("order_id")
    user_name := c.Query("user_name")

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

// 上传文件
func Upload(c *gin.Context) {
    file, header, err := c.Request.FormFile("file")
    if err != nil {
        c.String(http.StatusBadRequest, err.Error())
        log.Println("no such file: " + err.Error())
    }
    
    filename := header.Filename
    fmt.Println(file, err, filename)
    fmt.Printf("Type: %T %T %T\n", file, err, filename)

    // 创建一个文件，文件名为filename，这里的返回值out也是一个File指针
	out, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

    // 将file的内容拷贝到out
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	c.String(http.StatusCreated, "upload successful \n")
}