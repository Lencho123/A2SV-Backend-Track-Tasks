package data

import (
    "context"

    "github.com/Lencho123/A2SV-Backend-Track-Tasks/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

var TaskCollection *mongo.Collection // set in init

func CreateTask(ctx context.Context, task models.Task) (*mongo.InsertOneResult, error) {
    return TaskCollection.InsertOne(ctx, task)
}

func GetAllTasks(ctx context.Context) ([]models.Task, error) {
    var tasks []models.Task
    cursor, err := TaskCollection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    if err = cursor.All(ctx, &tasks); err != nil {
        return nil, err
    }
    return tasks, nil
}

func GetTaskByID(ctx context.Context, id primitive.ObjectID) (*models.Task, error) {
    var task models.Task
    err := TaskCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
    return &task, err
}

func UpdateTask(ctx context.Context, id primitive.ObjectID, task models.Task) error {
    _, err := TaskCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": task})
    return err
}

func DeleteTask(ctx context.Context, id primitive.ObjectID) error {
    _, err := TaskCollection.DeleteOne(ctx, bson.M{"_id": id})
    return err
}