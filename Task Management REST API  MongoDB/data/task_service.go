package data

import (
	"log"

	"github.com/Lencho123/A2SV-Backend-Track-Tasks/Task-Management-REST-API/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 1) DATABASE CONNECTIONS
var TaskCollection *mongo.Collection

func ConnectDB() {

	// ===========CONNECT TO DATABASE===============//
	// Setup the client configuration
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Failed to create Mongo client:", err)
	}

	// create context with timeout and cancel function to clear context resource by the end
	ctx, cancel := createContext()
	defer cancel()

	// Now connect the client to our server
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Important: Ping to confirm connection is alive
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB not responding:", err)
	}

	// ===========CREATE DB AND COLLECTION===============//
	// Create if not exist or access db and collection if exist
	TaskCollection = client.Database("taskdb").Collection("tasks")
	log.Println("✅ Connected to MongoDB and accessed collection.")

}

// 2) CRUD OPERATIONS ON THE DATABASE

// A) Get all tasks
func GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	ctx, cancel := createContext()
	defer cancel()

	// access cursor pointer that can iterete on tasks retrieved frem database
	cursor, err := TaskCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	// iterate over tasks using returned iterator
	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// Create new task and add to db collection
func AddNewTask(task models.Task) (*mongo.InsertOneResult, error) {
	ctx, cancel := createContext()
	defer cancel()
	result, err := TaskCollection.InsertOne(ctx, task)

	return result, err
}

// Retrieve task by ID
func GetTasksByID(ID primitive.ObjectID) (*models.Task, error) {
	ctx, cancel := createContext()
	defer cancel()

	var task models.Task
	err := TaskCollection.FindOne(ctx, bson.M{"_id": ID}).Decode(&task)

	if err != nil {
		return nil, err
	}
	return &task, nil
}


// Update task with given id

func UpdateTask(updatedTask *models.Task, ID string) error {
	ctx, cancel := createContext()
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(ID)

	if err !=nil{
		return err
	}

	TaskCollection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{
		"$set": bson.M{"title": updatedTask.Title, "description": updatedTask.Description},
	})

	return nil
}