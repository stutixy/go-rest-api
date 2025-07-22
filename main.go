package main

import (
	"go-rest-api/routes"

	"go-rest-api/controllers"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/service"

	"github.com/gin-gonic/gin"
)

func main() {
	database := db.ConnectDatabase()
	taskRepo := repository.NewTaskRepository(database)
	taskService := service.NewTaskService(taskRepo)
	taskController := controllers.NewTaskController(taskService)
	router := gin.Default()
	routes.RegisterTaskRoutes(router, taskController)
	router.Run(":8080")
}
