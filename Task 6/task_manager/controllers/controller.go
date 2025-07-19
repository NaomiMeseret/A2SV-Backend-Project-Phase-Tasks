package controllers

import (
	"fmt"
	"net/http"
	"task_manager/data"
	"task_manager/middleware"
	"task_manager/models"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *gin.Context){
	var u models.User
	if c.ShouldBindJSON(&u)!= nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error":"Invaild input"})
		return
	}
	user , err := data.CreateUser(u)
	if err != nil{
		c.IndentedJSON(http.StatusBadRequest , gin.H{"error" :err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated , gin.H{"id":user.ID.Hex(), "username":user.UserName, "role":user.Role})
}

func Login(c *gin.Context){
	var creds map[string]string
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	username := creds["username"]
	password := creds["password"]
	fmt.Printf("API Login - Username: %s, Password: %s\n", username, password)
	user, err := data.AuthenticateUser(username, password)
	if err !=nil{
		c.IndentedJSON(http.StatusUnauthorized , gin.H{"error":"Bad credentials"})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256 , jwt.MapClaims{
		"userID" :user.ID.Hex(),
		"role" :user.Role,
		"exp" : time.Now().Add(time.Hour *24).Unix(),
	})
	tokenString , _ := token.SignedString(middleware.JwtSecret)
	c.IndentedJSON(http.StatusOK, gin.H{"token" :tokenString})
}
func Promote (c *gin.Context){
	id := c.Param("id")
	if err := data.PromoteUser(id);err!=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"error" :err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK , gin.H{"message" :"Promoted to admin"})
}
func GetTask(c *gin.Context){
	tasks , err := data.GetTask()
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK  ,tasks)
}
func GetTaskByID(c *gin.Context){
	id:=c.Param("id")
	task , err:=data.GetTaskByID(id)
	if err!=nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}
func CreateTask(c *gin.Context){
	var newTask models.Task
	if err:=c.ShouldBindJSON(&newTask);err!=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":err.Error()})
		return
	}
	created ,err:=data.CreateTask(newTask)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated,created)
}
func UpdateTask(c *gin.Context){
	id:=c.Param("id")
	var updatedTask models.Task
	if err:=c.ShouldBindJSON(&updatedTask);err!=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":err.Error()})
		return
	}
	result,err := data.UpdateTask(id,updatedTask)
	if err!=nil{
		c.IndentedJSON(http.StatusNotFound ,gin.H{"message":err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}
func DeleteTask(c *gin.Context){
	id:=c.Param("id")
	if err:=data.DeleteTask(id);err!=nil{
		c.IndentedJSON(http.StatusNotFound , gin.H{"message":err.Error()})
	}
	c.Status(http.StatusNoContent)


}
func GetUsers(c *gin.Context) {
    users, err := data.GetAllUsers()
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    // Only return id, username, and role (not password)
    var result []gin.H
    for _, user := range users {
        result = append(result, gin.H{
            "id": user.ID.Hex(),
            "username": user.UserName,
            "role": user.Role,
        })
    }
    c.JSON(http.StatusOK, result)
}
