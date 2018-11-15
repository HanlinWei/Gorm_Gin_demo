package router

import (
    "github.com/gin-gonic/gin"
    handler "../handler"
    "net/http"
)

func InitRouter() *gin.Engine {
    router := gin.Default()

    router.GET("/list/", handler.ListAll)

    // The request responds to a url matching:  /add?amount=553.5&order_id=9&user_name=wabuguan&status=online&file_url=abc.txt
    router.GET("/add", handler.Store)

    router.GET("/update", handler.Update)

    router.GET("/delete", handler.Delete)

    router.GET("/fuzzysearch", handler.FuzzySearch)

    // curl -X POST http://localhost:8080/upload -F"file=@/home/qydev/image.jpg" -H "Content-Type: multipart/form-data"
    // curl -X POST http://localhost:8080/upload?order_id=1&user_name=Alice&file_url=image.jpg -F"file=@/home/qydev/image.jpg" -H "Content-Type: multipart/form-data"
    router.POST("/upload", handler.Upload)

    // router.Static("/db", "./db")
    
    router.StaticFS("/more_static", http.Dir("./db"))
    
    router.StaticFile("/export_excel.xlsx", handler.ExportExcel())
    
    router.StaticFile("/favicon.ico", "./resource/favicon.ico")

    return router
}