# 新手任务

0. 数据模型： type User struct {

   ​    ID       int64  `json:"id"`       // 列名为 `id`

   ​    Username string `json:"username"` // 列名为 `username`

   ​    Password string `json:"password"` // 列名为 `password`

   }

    1. 项目结构

       ```go
       ├── db   (数据库配置， GORM配置)
       │    ├── sqlite.go
       │    └── user.go
       ├── model
       │    └── model.go 
       ├── router  (Gin router)
       │    └── router.go 
       ├── service
       │    └── service.go 
       ├── handler  (Gin handler)
       ├── README.md
       └── main.go
       ```

