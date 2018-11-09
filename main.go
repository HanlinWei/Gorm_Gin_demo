package main

import (
    _ "./db"
    "./router"
)

func main() {
    router := router.InitRouter()
    router.Run()
}