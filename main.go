package main

import (
    _ "./db"
    "./router"
    service "./service"
)

func main() {
    router := router.InitRouter()
    service.ExportExcel()
    router.Run()
}