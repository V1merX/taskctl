package httpserver

import (
	"taskctl/internal/handler"
	"taskctl/internal/repository"
	"taskctl/internal/service"

	"github.com/gin-gonic/gin"
)

func (s *Server) setupRoutes() *gin.Engine {
	r := gin.New()

	taskRep := repository.NewTaskRepository(s.db)
	taskService := service.NewTaskService(taskRep)
	taskHandler := handler.NewTaskHandler(taskService)

	r.POST("/tasks", taskHandler.Create)
	r.GET("/tasks/:task_id", taskHandler.Get)
	r.DELETE("/tasks/:task_id", taskHandler.Delete)

	return r
}
