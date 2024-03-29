package router

import (
	"github.com/gin-gonic/gin"
	"github.com/maiyama18/tasks-gotemplate-vue/infra/database"
)

func Run() error {
	r := gin.Default()
	r.LoadHTMLGlob("infra/router/templates/*.html")

	taskInMemoryDatabase := database.NewTaskInMemoryDatabase()
	taskHandler := NewTaskHandler(taskInMemoryDatabase)

	r.GET("/", taskHandler.Index)

	r.POST("/api/tasks/new", taskHandler.AddTask)
	r.PUT("/api/tasks/toggle/:id", taskHandler.ToggleTask)
	r.DELETE("/api/tasks/delete/:id", taskHandler.DeleteTask)

	if err := r.Run(":8080"); err != nil {
		return err
	}
	return nil
}
