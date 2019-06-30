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

	if err := r.Run(":8080"); err != nil {
		return err
	}
	return nil
}
