package router

import (
	"net/http"

	"github.com/maiyama18/tasks-gotemplate-vue/domain/repository"

	"github.com/gin-gonic/gin"
	"github.com/maiyama18/tasks-gotemplate-vue/usecase"
)

type TaskHandler struct {
	taskUsecase *usecase.TaskUsecase
}

func NewTaskHandler(taskRepository repository.TaskRepository) *TaskHandler {
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	return &TaskHandler{taskUsecase: taskUsecase}
}

func (t *TaskHandler) Index(c *gin.Context) {
	tasks, err := t.taskUsecase.FindAll()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"tasks": tasks,
	})
}
