package main

import (
	"context"
	"log"
	"task_manager/Delivery/controllers"
	"task_manager/Delivery/routers"
	infrastructure "task_manager/Infrastructure"
	repositories "task_manager/Repositories"
	usecases "task_manager/Usecases"

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
	
	//Set Up Password Service
	passwordService :=infrastructure.NewPasswordService()
	//Set up usecases
	taskUsecase := usecases.NewTaskUsecase(taskRepo)
	userUsecase := usecases.NewUserUsecase(userRepo, passwordService)

	//Set up controllers
	taskController:=controllers.NewTaskController(taskUsecase)
	userController := controllers.NewUserController(userUsecase)

	//Set up Gin router and routes
	routers.InitRouter(taskController , userController)

	
}