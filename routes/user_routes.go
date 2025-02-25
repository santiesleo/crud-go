package routes

import (
    "gin-crud-project/controllers"
    "github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
    r.GET("/users", controllers.GetUsers)
    r.POST("/users", controllers.CreateUser)
    r.PATCH("/users/:id", controllers.UpdateUser)
    r.DELETE("/users/:id", controllers.DeleteUser)
}
