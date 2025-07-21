package routes

import (
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(router *gin.Engine) {
	taskGroup := router.Group("/task")
	{
		taskGroup.POST("", controllers.CreateTask)
		taskGroup.GET("", controllers.GetTasks)
		taskGroup.GET("/:id", controllers.GetTaskByID)
		taskGroup.PUT("/:id", controllers.UpdateTask)
		taskGroup.DELETE("/:id", controllers.DeleteTask)
	}
}
