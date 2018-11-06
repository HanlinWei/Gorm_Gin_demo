package main

import (
    _ "./db"
    "./router"
    orm "./db"
)

func main() {
    defer orm.Eloquent.Close()
    router := router.InitRouter()
    router.Run()
}