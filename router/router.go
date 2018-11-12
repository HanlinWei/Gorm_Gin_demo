package router

import (
    "github.com/gin-gonic/gin"
    handler "../handler"
)

func InitRouter() *gin.Engine {
    router := gin.Default()

    router.GET("/list/", handler.ListAll)

    router.GET("/add/:amount/:order_id/:user_name/:status/:file_url/", handler.Store)

    router.GET("/update/:amount/:order_id/:user_name/:status/:file_url/", handler.Update)

    router.GET("/delete/:order_id/:user_name/", handler.Delete)

    router.GET("/fuzzysearch/:keyword/:sortby/:desc/", handler.FuzzySearch)

    return router
}