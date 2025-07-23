package routers

import(
	"task_manager/Delivery/controllers"
	infrastructure "task_manager/Infrastructure"
	"github.com/gin-gonic/gin"
)
func BuildRoutes(r *gin.Engine , taskController *controllers.TaskController , userController *controllers.UserController){
	//Public routes
	r.POST("/register", userController.Register)
	r.POST("/login" , userController.Login)
	//Protected routes
	auth := r.Group("/")
	auth.Use(infrastructure.AuthRequired())
	{
		auth.POST("/tasks", taskController.CreateTask)
        auth.GET("/tasks", taskController.GetTasksByUserID)
        auth.GET("/tasks/:id", taskController.GetTaskByID)
        auth.PUT("/tasks/:id", taskController.UpdateTask)
        auth.DELETE("/tasks/:id", taskController.DeleteTask)
		//Admin- Only routes 
		auth.GET("/users", userController.GetAllUsers)
		auth.PUT("/promote/:id" , userController.PromoteUser)
	}

}