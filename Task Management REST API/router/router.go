package router

import (
	"github.com/Lencho123/A2SV-Backend-Track-Tasks/Task-Management-REST-API/controllers"
	"github.com/gin-gonic/gin"
)

func ServerRoutes() {
	var router = gin.Default()

	router.GET("/tasks", controllers.GetTasks)
	router.GET("/tasks/:id", controllers.GetTasksByID)
	router.PUT("/tasks/:id", controllers.PutTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)
	router.POST("/tasks", controllers.AddNewTask)

	router.Run("localhost:8080")
}
