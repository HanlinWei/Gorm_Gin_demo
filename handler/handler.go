package handler

import (
	"github.com/gin-gonic/gin"
	// model "../model"
	"io"
	"log"
	http "net/http"
	"os"
	"strconv"
	"time"

	service "../service"

	// "strings"
	"fmt"
)

// 列出所有数据
func ListAll(c *gin.Context) {
	x := "2017-02-27 17:30:20"
	p, _ := time.Parse(x, x)
	result := service.SelectByCreatedAt(p, time.Now())
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "列出所有数据",
		"data":    result,
	})
}

// 模糊查找
func FuzzySearch(c *gin.Context) {
	keyword := c.Query("keyword")
	sortby := c.Query("sortby")
	desc, err := strconv.ParseBool(c.Query("desc"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "参数desc不合法，应为true或false",
		})
		return
	}
	result, length, err := service.FuzzySearch(keyword, sortby, desc)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "查询失败",
		})
		return
	}
	if result == nil && length == 1 {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "参数sortby不合法，应为createdat或amount",
		})
		return
	}
	if length == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "未找到匹配",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":          1,
		"message":       "查询成功",
		"record number": length,
		"data":          result,
	})
}

// 添加数据
func Store(c *gin.Context) {
	service.BasicOperation(c, service.Insert, "创建")
}

// 修改数据
func Update(c *gin.Context) {
	service.BasicOperation(c, service.Change, "更新")
}

// 删除数据
func Delete(c *gin.Context) {
	order_id := c.Query("order_id")
	user_name := c.Query("user_name")

	failed := service.Delete(order_id, user_name)
	if failed != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":      1,
		"message":   "删除成功",
		"order_id":  order_id,
		"user_name": user_name,
	})
}

// 上传文件
func Upload(c *gin.Context) {

	amount, order_id, user_name, status, file_url, err := service.SolveParam(c)
	if len(order_id) > 0 && len(user_name) > 0 {
		if err != nil {
		}
		new_record := service.CreateDemoOrder(amount, order_id, user_name, status, file_url)
		failed := service.Change(new_record)
		if failed != nil {
			// c.String(http.StatusBadRequest, "Can't find the user.\n")
			return
		}
	}

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

func ExportExcel() string {
	return service.ExportExcel()
}
