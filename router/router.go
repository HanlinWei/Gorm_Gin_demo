package router

import (
    "github.com/gin-gonic/gin"
    service "../service"
)

func InitRouter() *gin.Engine {
    router := gin.Default()

    router.GET("/list/", service.ListAll)

    router.GET("/add/:amount/:order_id/:user_name/:status/:file_url/", service.Store)

    router.GET("/update/:amount/:order_id/:user_name/:status/:file_url/", service.Update)

    router.GET("/delete/:order_id/:user_name/", service.Delete)

    return router
}