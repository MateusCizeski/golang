package main

import (
	"github.com/gin-gonic/gin"
)

type Tasks struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

var taskList = []Tasks{
	{Id: 891, Title: "Estudar GO lang"},
	{Id: 123, Title: "Estudar"},
}

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", RouteTest)
	router.GET("/tarefas", GetAllTasks)
	router.POST("/tarefas", AddNewTask)
	router.GET("/tarefas/:id", GetTaskById)
	router.DELETE("/tarefas/:id", DeleteTaskById)
	router.PUT("/tarefas/:id", UpdateTaskById)
}
