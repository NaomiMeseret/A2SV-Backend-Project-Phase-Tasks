package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)
func GetTask(c *gin.Context){
	c.IndentedJSON(http.StatusOK,data.GetTask())
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
	created:=data.CreateTask(newTask)
	c.IndentedJSON(http.StatusCreated,created)
}
func UpdateTask(c *gin.Context){
	id:=c.Param("id")
	var updatedTask models.Task
	if err:=c.ShouldBindJSON(&updatedTask);err!=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":err.Error()})
		return
	}
	updatedTask,err := data.UpdateTask(id,updatedTask)
	if err!=nil{
		c.IndentedJSON(http.StatusNotFound ,gin.H{"message":err.Error()})
	}
	c.IndentedJSON(http.StatusOK, updatedTask)
}
func DeleteTask(c *gin.Context){
	id:=c.Param("id")
	if err:=data.DeleteTask(id);err!=nil{
		c.IndentedJSON(http.StatusNotFound , gin.H{"message":err.Error()})
	}
	c.Status(http.StatusNoContent)


}
