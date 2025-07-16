package main

import (
	"task_manager/router"
)
func main(){
	server:=router.BuildRouter()
	server.Run(":8080")
}