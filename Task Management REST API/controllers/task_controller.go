package controllers

import (
	"github.com/Lencho123/A2SV-Backend-Track-Tasks/Task-Management-REST-API/data"
	"github.com/Lencho123/A2SV-Backend-Track-Tasks/Task-Management-REST-API/models"

	// "github.com/Lencho123/A2SV-Backend-Track-Tasks/Task-Management-REST-API/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	// "fmt"
)

// FUNCTION TO GET ALL TASKS AVAILABLE
func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, data.TaskDatas)
}

// RETRIEVE SPECIFIC TASK BASED ON ITS ID
func GetTasksByID(c *gin.Context) {
	ID := c.Param("id")
	IntID, err := strconv.Atoi(ID)

	// fmt.Println(IntID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid ID"})
		return
	}
	task := data.GetTaskWithID(IntID)

	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Something went wrong"})
		// map[string]interface{}{"Error": "Something went wrong"} === gin.H{"Error": "Something went wrong"}
		return
	}

	c.JSON(http.StatusOK, task)
}

// UPDATE TASK WITH SOME ID

func PutTask(c *gin.Context) {
	var UpdatedTask models.Task
	bindErr := c.BindJSON(&UpdatedTask)
	ID := c.Param("id")

	IntID, err := strconv.Atoi(ID)

	if err != nil || bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Input"})
		return
	}

	message := data.PutTask(IntID, UpdatedTask)

	if message == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Couldn't find data with this id"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Seccuss": "Data got updated successfully"})
}

// DELETE TASK WITH SPECIFIC ID
func DeleteTask(c *gin.Context) {
	ID := c.Param("id")
	IntID, err := strconv.Atoi(ID)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	status := data.DeleteTask(IntID)

	if status == 1{
		c.JSON(http.StatusOK, gin.H{"Success":"Task deleted seccusfully"})
	}else{
	c.JSON(http.StatusNotFound, gin.H{"Error":"Task with that id couldn't be found"})
	}
}

// ADD NEW TASK
func AddNewTask(c *gin.Context){
	var NewTask models.Task
	bindErr := c.BindJSON(&NewTask)

	if bindErr != nil{
		c.JSON(http.StatusBadRequest, gin.H{"Error":"Invalid Task"})
		return
	}
	data.TaskDatas = append(data.TaskDatas, NewTask)
	c.JSON(http.StatusOK, gin.H{"Succus":"Taks added succesfully"})
}