package router

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)
func BuildRouter()*gin.Engine{
	r:=gin.Default()
	// Public endpoints
	r.POST("/register" , controllers.Register)
	r.POST("/login" , controllers.Login)
	
	// Protected:any vaild user
	auth := r.Group("/", middleware.AuthRequired())
	auth.GET("/tasks",controllers.GetTask)
	auth.GET("/tasks/:id",controllers.GetTaskByID)

	// Admin only for write actions
	admin:=auth.Group("/", middleware.AdminOnly())
	admin.POST("/tasks",controllers.CreateTask)
	admin.PUT("/tasks/:id", controllers.UpdateTask)
	admin.DELETE("/tasks/:id", controllers.DeleteTask)
	admin.PUT("/promote/:id" , controllers.Promote)
	admin.GET("/users", controllers.GetUsers) 
	return r

}