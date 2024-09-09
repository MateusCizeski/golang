package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouteTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Primeira api com GO rodando!",
	})
}

// buscando todas as tarefas
func GetAllTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, taskList)
}

// cadastro novas tarefas
func AddNewTask(c *gin.Context) {
	var newTask Tasks

	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newTask.Id = len(taskList) + 1
	taskList = append(taskList, newTask)
	c.JSON(http.StatusOK, newTask)
}

// buscando tarefa por id
func GetTaskById(c *gin.Context) {
	id := c.Param("id")

	for _, task := range taskList {
		if fmt.Sprintf("%d", task.Id) == id {
			c.JSON(http.StatusOK, task)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Tarefa não foi encontrada",
	})
}

// deletando tarefa por id
func DeleteTaskById(c *gin.Context) {
	id := c.Param("id")

	for index, task := range taskList {
		if fmt.Sprintf("%d", task.Id) == id {
			taskList = append(taskList[:index], taskList[index+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Tarefa deletada com sucesso."})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Tarefa não encontrada."})
}

// editar tarefa
func UpdateTaskById(c *gin.Context) {
	id := c.Param("id")
	var updatedTask Tasks

	if err := c.BindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Terafa não encontrada."})
		return
	}

	for index, task := range taskList {
		if fmt.Sprintf("%d", task.Id) == id {
			updatedTask.Id = task.Id
			taskList[index] = updatedTask
			c.JSON(http.StatusOK, gin.H{"message": "Tarefa atualziada com sucesso."})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Tarefa não encontrada."})
}
