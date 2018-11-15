# 新手任务

0. 数据模型： type User struct {

​    ID       int64  `json:"id"`       // 列名为 `id`

​    Username string `json:"username"` // 列名为 `username`

​    Password string `json:"password"` // 列名为 `password`

​    }

 1. 项目结构

    ```go
    ├── db   (数据库配置， GORM配置)
    │    └── database.go
    ├── model
    │    └── model.go 
    ├── router  (Gin router)
    │    └── router.go 
    ├── service
    │    └── service.go 
    ├── handler  (Gin handler)
    │    └── handler.go
    ├── README.md
    └── main.go
    ```

2. 实现 4个需求：
   ​                        1） 创建 demo_order
   ​                        2） 更新 demo_order （amount、status、file_url）
   ​                        3） 获取 demo_order 详情
   ​                        4） 获取 demo_order 列表 （需要包含： 模糊查找、根据创建时间，金额排序）



3. 根据需求编写 Restful Api接口
4. service 下代码需要编写测试
5. 需要实现一个 Gorm 事物 SQL
6. 需要实现一个Gin 文件上传的功能， 并且更新 demo_order ： file_url
7. 同时实现Gin 文件下载
8. 使用 https://github.com/tealeg/xlsx 这个库 将demo_order 所有数据以excel形式导出来(可以下载)