package controllers

import (
	"go-rest-api/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if strings.TrimSpace(newTask.Title) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
	}
	newTask.ID = models.NextID
	models.NextID++
	models.Tasks = append(models.Tasks, newTask)
	c.JSON(http.StatusCreated, newTask)
}

func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, models.Tasks)
}

func GetTaskByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task"})
		return
	}
	for _, task := range models.Tasks {
		if task.ID == id {
			c.JSON(http.StatusOK, task)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
}

func UpdateTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task"})
		return
	}
	var updateTask models.Task

	if err := c.BindJSON(&updateTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, task := range models.Tasks {
		if task.ID == id {
			updateTask.ID = id
			models.Tasks[i] = updateTask
			c.JSON(http.StatusOK, models.Tasks[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
}

func DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range models.Tasks {
		if task.ID == id {
			models.Tasks = append(models.Tasks[:i], models.Tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
}
