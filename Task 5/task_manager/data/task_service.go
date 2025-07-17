package data

import (
	"task_manager/models"
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"fmt"
)
var taskCollection *mongo.Collection
func InitMongo(){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions:=options.Client().ApplyURI("mongodb://localhost:27017")
	client , err:=mongo.Connect(ctx, clientOptions)
	if err !=nil{
		log.Fatal("Mongo Connect error:err")
	}
	err = client.Ping(ctx, nil)
	if err!=nil{
		log.Fatal("MongoPing error:",err)
	}
	taskCollection  = client.Database("taskdb").Collection("tasks")
	fmt.Println(" ✅ Connected to MongoDB")
}

func GetTask()([]models.Task, error){
	cursor , err := taskCollection.Find(context.TODO(),bson.M{})
	if err!=nil{
		return nil , err
	}
	defer cursor.Close(context.TODO())
	var tasks []models.Task
	for cursor.Next(context.TODO()){
		var task models.Task
		if err:=cursor.Decode(&task);err!=nil{
			return nil , err
		}
		tasks = append(tasks, task)
	}
	return tasks , nil
}
func GetTaskByID(id string) (models.Task ,error){
	var task models.Task
	objID, err := primitive.ObjectIDFromHex(id)
	if err!=nil{
		return task ,err
	}
	filter := bson.M{"_id":objID}
	result := taskCollection.FindOne(context.TODO(),filter)
	err =result.Decode(&task)
	return task , err

}
func CreateTask(task models.Task) (models.Task , error){
	result, err:= taskCollection.InsertOne(context.TODO() , task)
	if err!=nil{
		return task , err
	}
	objID , ok := result.InsertedID.(primitive.ObjectID)
	if !ok{
		return task , fmt.Errorf("can't convert InsertedID")
	}
	task.ID = objID
	return task ,nil
}
func UpdateTask(id string ,updatedTask models.Task) (models.Task, error){
	objID , err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return updatedTask ,err
	}
	filter := bson.M{"_id" : objID}
	update :=bson.M{"$set":updatedTask}
	_ ,err = taskCollection.UpdateOne(context.TODO() ,filter ,update)
	if err != nil {
		return updatedTask , err
	}
	updatedTask.ID = objID
	return updatedTask ,nil

}
func DeleteTask(id string)error{
	objID , err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return err
	}
	_ , err  = taskCollection.DeleteOne(context.TODO() , bson.M{"_id":objID})
	return err
}
