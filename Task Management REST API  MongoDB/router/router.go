package router

import (
	"github.com/Lencho123/A2SV-Backend-Track-Tasks/Task-Management-REST-API/controllers"
	"github.com/gin-gonic/gin"
)

func ServerRoutes() {
	router := gin.Default()

	router.GET("/tasks", controllers.GetTasks)
	router.POST("/tasks", controllers.AddNewTask)
	router.GET("/tasks/:id", controllers.GetTaskByID)
	router.PUT("/tasks/:id", controllers.UpdateTask)
	// router.DELETE("/tasks/:id", controllers.DeleteTask)

	router.Run("localhost:8080")
	}