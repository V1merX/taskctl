package httpserver

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) setupRoutes() *gin.Engine {
	var r gin.Engine

	// taskRep	:= storage.NewStorage()
	// taskService := service.NewTaskService()
	// taskHandler := handler.NewTaskHandler(taskService)

	// TODO: add routes
	// r.POST("/tasks", createTaskHandler)
	// r.GET("/tasks/:id", getTaskHandler)
	// r.DELETE("/tasks/:id", deleteTaskHandler)

	return &r
}
