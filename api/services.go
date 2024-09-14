package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Tasks struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func RouteTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Primeira api com GO rodando!",
	})
}

// buscando todas as tarefas
func GetAllTasks(c *gin.Context) {
	rows, err := DB.Query("SELECT id, title from tasks")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	defer rows.Close()

	var tasks []Tasks

	for rows.Next() {
		var task Tasks
		if err := rows.Scan(&task.Id, &task.Title); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, tasks)
}

// cadastro novas tarefas
func AddNewTask(c *gin.Context) {
	var newTask Tasks

	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := DB.Exec("INSERT INTO tasks (title) VALUES (?)", newTask.Title)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := result.LastInsertId()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newTask.Id = int(id)

	c.JSON(http.StatusCreated, newTask)
}

// buscando tarefa por id
func GetTaskById(c *gin.Context) {
	id := c.Param("id")

	var task Tasks

	row := DB.QueryRow("SELECT id, title FROM tasks WHERE id = ?", id)

	if err := row.Scan(&task.Id, &task.Title); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

// deletando tarefa por id
func DeleteTaskById(c *gin.Context) {
	id := c.Param("id")

	_, err := DB.Exec("DELETE FROM tasks WHERE id = ?", id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tarefa deletada com sucesso."})
}

// editar tarefa
func UpdateTaskById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var updatedTask Tasks

	if err := c.BindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Terafa não encontrada."})
		return
	}

	_, err := DB.Exec("UPDATE tasks SET title = ? WHERE id = ?", updatedTask.Title, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	updatedTask.Id = id
	c.JSON(http.StatusOK, updatedTask)
}
