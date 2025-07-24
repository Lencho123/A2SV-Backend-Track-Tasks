package controllers

import (
	"net/http"

	"github.com/Lencho123/A2SV-Backend-Track-Tasks/Task-Management-REST-API/data"
	"github.com/Lencho123/A2SV-Backend-Track-Tasks/Task-Management-REST-API/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get All tasks from data collection
func GetTasks(c *gin.Context) {
	tasks, err := data.GetTasks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// Add new task to data collection
func AddNewTask(c *gin.Context) {
	var task models.Task

	err := c.BindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result, err := data.AddNewTask(task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}
	c.JSON(http.StatusCreated, result)
}

// Retrive task by its ID
func GetTaskByID(c *gin.Context) {
	strID := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(strID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid ID format!"})
		return
	}

	task, err := data.GetTasksByID(objID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Something wents wrong"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// Update existing taks by id
func UpdateTask(c *gin.Context){
	var updatedTask models.Task

	ID := c.Param("id")


	err := c.BindJSON(&updatedTask)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"Error":"Invalid input!"})
		return
	}

	err = data.UpdateTask(&updatedTask, ID)
}