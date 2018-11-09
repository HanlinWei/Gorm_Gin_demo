package router

import (
    "github.com/gin-gonic/gin"
    service "../service"
)

func InitRouter() *gin.Engine {
    router := gin.Default()

    router.GET("/list", service.ListAll)

    router.POST("/add", service.Store)

    router.PUT("/Update/:id", service.Update)

    router.DELETE("/Destroy/:id", service.Destroy)

    return router
}