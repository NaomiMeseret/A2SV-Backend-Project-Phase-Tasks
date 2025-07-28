package repositories

import (
	"context"
	"errors"
	domain "task_manager/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//mongoTaskReposistory is the MongoDB implemntation of the TaskRepository interface
type mongoTaskRepository struct{
	collection *mongo.Collection
}

//NewMongoTaskReposistory creates a new TaskRepository with the given MongoDB collection.
func NewMongoTaskReposistory(collection *mongo.Collection)domain.ITaskRepository{
	return &mongoTaskRepository{collection: collection}
}

func (r *mongoTaskRepository)CreateTask(task *domain.Task)error{
	taskData:=bson.M{
		"title": task.Title , 
		"description": task.Description,
		"date": task.Date,
		"status": task.Status,
		"user_id": task.UserID,
	}
	result , err:=r.collection.InsertOne(context.TODO(),taskData)
	if err != nil{
		return err
	}
	objID , ok:=result.InsertedID.(primitive.ObjectID)
	if ok{
		task.ID = objID.Hex()
	}
	return nil

}

func (r *mongoTaskRepository) GetTaskByID(id string) (*domain.Task, error) {
    oid, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, errors.New("invalid task ID")
    }
    filter := bson.M{"_id": oid}
    var taskData bson.M
    err = r.collection.FindOne(context.TODO(), filter).Decode(&taskData)
    if err != nil {
        return nil, err
    }
    task := &domain.Task{
        ID:          id,
        Title:       taskData["title"].(string),
        Description: taskData["description"].(string),
        Date:        taskData["date"].(string),
        Status:      taskData["status"].(string),
        UserID:      taskData["user_id"].(string),
    }
    return task, nil
}

func (r *mongoTaskRepository) GetTasksByUserID(userID string) ([]*domain.Task, error) {
    filter := bson.M{"user_id": userID}
    cursor, err := r.collection.Find(context.TODO(), filter)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())

    var tasks []*domain.Task
    for cursor.Next(context.TODO()) {
        var taskData bson.M
        if err := cursor.Decode(&taskData); err != nil {
            continue
        }
        id := ""
        if objID, ok := taskData["_id"].(primitive.ObjectID); ok {
            id = objID.Hex()
        }
        task := &domain.Task{
            ID: id,
            Title: taskData["title"].(string),
            Description: taskData["description"].(string),
            Date: taskData["date"].(string),
            Status: taskData["status"].(string),
            UserID: taskData["user_id"].(string),
        }
        tasks = append(tasks, task)
    }
    return tasks, nil
}
func (r *mongoTaskRepository) UpdateTask(task *domain.Task) error {
    objID, err := primitive.ObjectIDFromHex(task.ID)
    if err != nil {
        return errors.New("invalid task ID")
    }
    filter := bson.M{"_id": objID}
    update := bson.M{
        "$set": bson.M{
            "title":task.Title,
            "description": task.Description,
            "date": task.Date,
            "status": task.Status,
            "user_id": task.UserID,
        },
    }
    _, err = r.collection.UpdateOne(context.TODO(), filter, update)
    return err
}
func (r *mongoTaskRepository) DeleteTask(id string) error {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return errors.New("invalid task ID")
    }
    filter := bson.M{"_id": objID}
    _, err = r.collection.DeleteOne(context.TODO(), filter)
    return err
} 
