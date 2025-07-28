package router

import (
    "github.com/Lencho123/A2SV-Backend-Track-Tasks/controllers"
    "github.com/Lencho123/A2SV-Backend-Track-Tasks/middleware"
    "github.com/gin-gonic/gin"
)

func SetupAndRun() {
    r := gin.Default()
    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)

    auth := r.Group("/")
    auth.Use(middleware.AuthMiddleware())
    {
        auth.GET("/tasks", controllers.GetAllTasks)        // All users
        auth.GET("/tasks/:id", controllers.GetTaskByID)     // All users
        auth.POST("/tasks", middleware.AdminOnly(), controllers.CreateTask)
        auth.PUT("/tasks/:id", middleware.AdminOnly(), controllers.UpdateTask)
        auth.DELETE("/tasks/:id", middleware.AdminOnly(), controllers.DeleteTask)
        auth.POST("/promote/:username", middleware.AdminOnly(), controllers.PromoteUser)
    }

    r.Run(":8080")
}