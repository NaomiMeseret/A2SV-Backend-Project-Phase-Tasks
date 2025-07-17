package main

import (
	"task_manager/router"
	"task_manager/data"
)
func main(){
	data.InitMongo()
	server:=router.BuildRouter()
	server.Run(":8080")
}