package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"
	"github.com/gin-gonic/gin"
)
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
