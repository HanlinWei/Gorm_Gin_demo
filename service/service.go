package service

import (
    "github.com/gin-gonic/gin"
	model "../model"
	orm "../db"
    http "net/http"
    "strconv"
)

// 列出所有数据
func ListAll(c *gin.Context) {
    // var user model.User
    // user.Username = c.Request.FormValue("username")
    // user.Password = c.Request.FormValue("password")
    result, err := AllUsers()

    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "code":    -1,
            "message": "抱歉未找到相关信息",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "code": 1,
        "data": result,
    })
}

//添加数据
func Store(c *gin.Context) {
    var user model.User
    user.Username = c.Request.FormValue("username")
    user.Password = c.Request.FormValue("password")
    id, err := Insert(user)

    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "code":    -1,
            "message": "添加失败",
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "code":  1,
        "message": "添加成功",
        "data":    id,
    })
}

//修改数据
func Update(c *gin.Context) {
    var user model.User
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    user.Password = c.Request.FormValue("password")
    result, err := Change(id)
    if err != nil || result.ID == 0 {
        c.JSON(http.StatusNotFound, gin.H{
            "code":    -1,
            "message": "修改失败",
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "code":  1,
        "message": "修改成功",
    })
}

//删除数据
func Destroy(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    result, err := Delete(id)
    if err != nil || result.ID == 0 {
        c.JSON(http.StatusNotFound, gin.H{
            "code":    -1,
            "message": "删除失败",
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "code":  1,
        "message": "删除成功",
    })
}

//添加
func Insert(user model.User) (id int64, err error) {

    //添加数据
    result := orm.Eloquent.Create(&user)
    id =user.ID
    if result.Error != nil {
        err = result.Error
        return -1, nil
    }
    return id, nil
}

//列表
func AllUsers() (users model.User, err error) {
    var find = orm.Eloquent.Where("id >= ", 0).Find(&users)
    if err = find.Error; err != nil {
        return users, err
    }
    return users, nil
}

//修改
func Change(id int64) (updateUser model.User, err error) {
	err = orm.Eloquent.Select([]string{"id", "username"}).First(&updateUser, id).Error
    if err != nil {
        return
    }

    //参数1:是要修改的数据
	//参数2:是修改的数据
	err = orm.Eloquent.Model(&updateUser).Updates(&updateUser).Error
    if err != nil {
        return
    }
    return
}

//删除数据
func Delete(id int64) (Result model.User, err error) {
    if err = orm.Eloquent.Select([]string{"id"}).First(&Result, id).Error; err != nil {
        return
    }

    if err = orm.Eloquent.Delete(&Result).Error; err != nil {
        return
    }
    return
}