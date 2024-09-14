package main

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", RouteTest)
	router.GET("/tarefas", GetAllTasks)
	router.POST("/tarefas", AddNewTask)
	router.GET("/tarefas/:id", GetTaskById)
	router.DELETE("/tarefas/:id", DeleteTaskById)
	router.PUT("/tarefas/:id", UpdateTaskById)
}
