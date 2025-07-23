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

func (uc *UserController) Register(c *gin.Context){
	var user domain.User
	err:=c.ShouldBindJSON(&user)
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest , gin.H{"error":"Invalid input"})
		return 
	}
	hashed , err := infrastructure.HashPassword(user.Password)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error":"Failed to hash password"})
		return
	}
	user.Password = hashed
	err = uc.userUsecase.RegisterUser(&user)
	if err != nil{
		c.IndentedJSON(http.StatusBadRequest , gin.H{"error": err.Error()})
		return 
	}
	c.IndentedJSON(http.StatusOK , gin.H{"message":"User registered successfully"})

}

func (uc *UserController) Login (c *gin.Context){
	var input struct{
		UserName string `json:"username"`
		Password string `json:"password"`
	}
	err:=c.ShouldBindJSON(&input)
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest , gin.H{"error":"Invalid input"})
		return 
	}
	user , err := uc.userUsecase.LoginUser(input.UserName,input.Password)
	if err != nil{
		c.IndentedJSON(http.StatusUnauthorized ,gin.H{"error":"Invalid credentials"})
	}
	// Check Password
	if !infrastructure.CheckPasswordHash(input.Password , user.Password){
		c.IndentedJSON(http.StatusUnauthorized , gin.H{"error":"Invalid credentials"})
		return
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

func (tc *TaskController) CreateTask (c *gin.Context){
	var task domain.Task
	err:=c.ShouldBindJSON(&task)
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest , gin.H{"error":"Invalid input"})
		return
	}
	//Get user_id from the context (set by asuth middleware)
	userID , ok :=c.Get("user_id")
	if !ok {
		c.IndentedJSON(http.StatusUnauthorized , gin.H{"error":"User not authenticated"})
		return 
	}
	task.UserID = userID.(string)
	err= tc.taskUsecase.CreateTask(&task)
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

func (tc *TaskController)UpdateTask(c *gin.Context){
	id := c.Param("id")
    var task domain.Task
    err := c.ShouldBindJSON(&task)
	if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    task.ID = id
    err = tc.taskUsecase.UpdateTask(&task)
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.IndentedJSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
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