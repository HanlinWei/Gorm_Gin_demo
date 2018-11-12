package router

import (
    "github.com/gin-gonic/gin"
    handler "../handler"
)

func InitRouter() *gin.Engine {
    router := gin.Default()

    router.GET("/list/", handler.ListAll)

    // The request responds to a url matching:  /add?amount=553.5&order_id=9&user_name=wabuguan&status=online&file_url=abc.txt
    router.GET("/add", handler.Store)

    router.GET("/update", handler.Update)

    router.GET("/delete", handler.Delete)

    router.GET("/fuzzysearch", handler.FuzzySearch)

    router.POST("/upload", handler.Upload)

    return router
}