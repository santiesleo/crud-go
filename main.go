package main

import (
    "gin-crud-project/db"
    "gin-crud-project/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    db.Connect()

    routes.UserRoutes(r)

    r.Run(":8080")
}
