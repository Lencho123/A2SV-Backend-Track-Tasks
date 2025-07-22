package data

import (
	"github.com/Lencho123/A2SV-Backend-Track-Tasks/Task-Management-REST-API/models"
)

var TaskDatas = []models.Task{
	{ID: 1, Title: "Read Clean Code", Description: "Go through Chapter 3: Functions"},
	{ID: 2, Title: "Refactor Service Layer", Description: "Split business logic from controller"},
	{ID: 3, Title: "Setup Database", Description: "Integrate SQLite or PostgreSQL with GORM"},
	{ID: 4, Title: "Test Endpoints", Description: "Use Postman to test all API routes"},
	{ID: 5, Title: "Deploy to Render", Description: "Push the app live using Docker and Render"},
}

func GetTaskWithID(ID int) *models.Task {
	for index := range TaskDatas {
		if TaskDatas[index].ID == ID {
			return &TaskDatas[index]
		}
	}
	return nil
}

// Update task

func PutTask(ID int, NewTask models.Task) int {
	for index := range TaskDatas {
		if TaskDatas[index].ID == ID {
			p := &TaskDatas[index]
			p.Title = NewTask.Title
			p.Description = NewTask.Description

			return 1
		}
	}

	return 0
}

// Delete task with id = ID

func DeleteTask(ID int) int {
	for index := range TaskDatas{
		if TaskDatas[index].ID == ID {
			TaskDatas = append(TaskDatas[:index],TaskDatas[index+1:]... )
			return 1
		}
	}
	return 0
}