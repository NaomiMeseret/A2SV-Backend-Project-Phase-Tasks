package main

import (
	"context"
	"log"
	"task_manager/Delivery/controllers"
	"task_manager/Delivery/routers"
	repositories "task_manager/Repositories"
	usecases "task_manager/Usecases"
	
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func main(){
	client, err:=mongo.Connect(context.Background() , options.Client().ApplyURI("mongodb://localhost:27017"))
	if err!=nil{
		log.Fatal("Failed to connect to MongoDB: ", err)
	}
	log.Println("✅Connected to MongoDB")
	db:=client.Database("task_manager")
	//Set up repositories
	taskCollection :=db.Collection("tasks")
	userCollection :=db.Collection("users")
	taskRepo:=repositories.NewMongoTaskReposistory(taskCollection)
	userRepo := repositories.NewMongoUserReposistory(userCollection)
	
	//Set up usecases
	taskUsecase := usecases.NewTaskUsecase(taskRepo)
	userUsecase := usecases.NewUserUsecase(userRepo)

	//Set up controllers
	taskController:=controllers.NewTaskController(taskUsecase)
	userController := controllers.NewUserController(userUsecase)

	//Set up Gin router and routes
	r:=gin.Default()
	routers.BuildRoutes(r, taskController , userController)

	//Start the server
	err= r.Run(":8080")
	if err!=nil{
		log.Fatal("Failed to start server: ", err)
	}
}