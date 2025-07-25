package controllers

import (
	"net/http"
	domain "task_manager/Domain"
	infrastructure "task_manager/Infrastructure"

	"github.com/gin-gonic/gin"
)

//UserController handles user-related HTTP requests.
type UserController struct{
	userUsecase domain.UserUsecase
}

//NewUserController creates new User Controller
func NewUserController (userUsecase domain.UserUsecase) *UserController{
	return &UserController{userUsecase:userUsecase}
}

type UserDTO struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
func (uc *UserController) ChangeToDomain(userDTO *UserDTO) *domain.User{
	var user domain.User
	user.Email = userDTO.Email
	user.Password = userDTO.Password
	return &user
}

func (uc *UserController) Register(c *gin.Context){
	var userDTO UserDTO
	err:=c.ShouldBindJSON(&userDTO)
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest , gin.H{"error":"Invalid input"})
		return 
	}
	user :=uc.ChangeToDomain(&userDTO)
	err =uc.userUsecase.RegisterUser(user)
	if err != nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK , gin.H{"message":"User registered successfully"})

}

func (uc *UserController) Login (c *gin.Context){
	var userDTO UserDTO
	err:=c.ShouldBindJSON(&userDTO)
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest , gin.H{"error":"Invalid input"})
		return 
	}
	user , err := uc.userUsecase.LoginUser(userDTO.Email,userDTO.Password)
	if err != nil{
		c.IndentedJSON(http.StatusUnauthorized ,gin.H{"error":"Invalid credentials"})
	}
	//Generate JWT token
	token , err := infrastructure.GenerateJWT(user.ID , user.Role)
	if err!=nil{
		c.IndentedJSON(http.StatusInternalServerError , gin.H{"error":"Failed to generate token"})
		return 
	}
	c.IndentedJSON(http.StatusOK ,gin.H{"token":token})

}

func (uc *UserController) GetAllUsers (c *gin.Context){
	role , _:=c.Get("role")
	if role != "admin"{
		c.IndentedJSON(http.StatusForbidden ,gin.H{"error":"Admins only"})
		return
	}
	users, err := uc.userUsecase.GetAllUsers()
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK , users)
}
func (uc *UserController) PromoteUser (c *gin.Context){
	role , _:=c.Get("role")
	if role != "admin"{
		c.IndentedJSON(http.StatusForbidden ,gin.H{"error":"Admins only"})
		return
	}
	id :=c.Param("id")
	err:=uc.userUsecase.PromoteUser(id)
	if err != nil{
		c.IndentedJSON(http.StatusBadRequest , gin.H{"error":err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK , gin.H{"message":"Promoted to admin"})
}

//TaskController handles task-related HTTP requests
type TaskController struct{
	taskUsecase domain.TaskUsecase
}

//NewTaskController creates a new TaskController
func NewTaskController(taskUsecase domain.TaskUsecase)*TaskController{
	return &TaskController{taskUsecase:taskUsecase}
}
type TaskDTO struct{
	Title string `json:"title"`
	Description string `json:"description"`
	Date string `json:"date"`
	Status string `json:"status"`
}

func (tc *TaskController)ChangeTaskToDomain(taskDTO *TaskDTO , userID string)*domain.Task{
	var task domain.Task
	task.Title  = taskDTO.Title
	task.Description = taskDTO.Description
	task.Date = taskDTO.Date
	task.Status = taskDTO.Status
	task.UserID = userID
	return &task
}
func (tc *TaskController) CreateTask (c *gin.Context){
	var taskDTO TaskDTO
	err:=c.ShouldBindJSON(&taskDTO)
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest , gin.H{"error":"Invalid input"})
		return
	}
	//Get user_id from the context (set by auth middleware)
	userID , ok :=c.Get("user_id")
	if !ok {
		c.IndentedJSON(http.StatusUnauthorized , gin.H{"error":"User not authenticated"})
		return 
	}
	task :=tc.ChangeTaskToDomain(&taskDTO , userID.(string))
	err= tc.taskUsecase.CreateTask(task)
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK , gin.H{"message":"Task created successfully"})
}
func (tc *TaskController)GetTasksByUserID(c *gin.Context){
	userID , ok := c.Get("user_id")
	if !ok {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
        return
	}
	tasks, err := tc.taskUsecase.GetTasksByUserID(userID.(string))
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.IndentedJSON(http.StatusOK, tasks)
}

func (tc *TaskController)GetTaskByID (c *gin.Context){
	id := c.Param("id")
    task, err := tc.taskUsecase.GetTaskByID(id)
    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }
    c.IndentedJSON(http.StatusOK, task)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
    id := c.Param("id")
    var taskDTO TaskDTO
    if err := c.ShouldBindJSON(&taskDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    // Fetch the existing task to keep the original userID
    existingTask, err := tc.taskUsecase.GetTaskByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }
    // Only update allowed fields
    existingTask.Title = taskDTO.Title
    existingTask.Description = taskDTO.Description
    existingTask.Date = taskDTO.Date
    existingTask.Status = taskDTO.Status
    err = tc.taskUsecase.UpdateTask(existingTask)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}
func (tc *TaskController)DeleteTask(c *gin.Context){
	id:=c.Param("id")
	err:=tc.taskUsecase.DeleteTask(id)
	if err != nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return 
	}
	c.IndentedJSON(http.StatusOK , gin.H{"message":"Task deleted successfully"})

}