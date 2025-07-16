package data

import (
	"task_manager/models"
	"errors"
)
var tasks = []models.Task{
	{ID:"1",
	Title:"Note-Taking",
	Description:"Take a note for a daily standup meeting", 
	Date:"2025-07-16",
	Status:"Completed",
},
{ID:"2",
	Title:"Do Task 3",
	Description:"Read documenation and complete the task", 
	Date:"2025-07-16",
	Status:"In Progress",
},

}
func GetTask()[]models.Task{
	return tasks
}
func GetTaskByID(id string) (models.Task ,error){
	for _ , task :=range tasks{
		if task.ID == id{
			return task , nil
		}
	}
	return models.Task{}, errors.New("task not found")

}
func CreateTask(task models.Task) models.Task{
	tasks = append(tasks, task)
	return task
}
func UpdateTask(id string ,updated models.Task) (models.Task, error){
	for i ,t:=range tasks{
		if t.ID ==id{
			tasks[i] =updated
			return updated ,nil	

		}
	}
	return models.Task{} , errors.New("task not found")
}
func DeleteTask(id string)error{
	for i , t:=range tasks{
		if t.ID == id{
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
