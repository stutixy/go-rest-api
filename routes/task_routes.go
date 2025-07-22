package routes

import (
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(router *gin.Engine, controller *controllers.TaskController) {
	taskGroup := router.Group("/task")
	{
		taskGroup.POST("", controller.CreateTask)
		taskGroup.GET("", controller.GetTasks)
		taskGroup.GET("/:id", controller.GetTaskByID)
		taskGroup.PUT("/:id", controller.UpdateTask)
		taskGroup.DELETE("/:id", controller.DeleteTask)
	}
}
