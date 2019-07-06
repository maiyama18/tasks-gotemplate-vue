package router

import (
	"net/http"
	"strconv"

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

	var taskResponses []*TaskResponse
	for _, t := range tasks {
		taskResponses = append(taskResponses, NewTaskResponse(t))
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"tasks": taskResponses,
	})
}

func (t *TaskHandler) AddTask(c *gin.Context) {
	var addTaskRequest AddTaskRequest
	if err := c.BindJSON(&addTaskRequest); err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
	}

	task, err := t.taskUsecase.Add(addTaskRequest.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
	}

	c.JSON(http.StatusOK, NewTaskResponse(task))
}

func (t *TaskHandler) ToggleTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err.Error()))
	}

	task, err := t.taskUsecase.Toggle(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewErrorResponse(err.Error()))
	}

	c.JSON(http.StatusOK, NewTaskResponse(task))
}
