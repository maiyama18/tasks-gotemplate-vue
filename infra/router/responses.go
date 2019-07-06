package router

import "github.com/maiyama18/tasks-gotemplate-vue/domain/model"

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
	}
}

type TaskResponse struct {
	ID        uint64 `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func NewTaskResponse(task *model.Task) *TaskResponse {
	return &TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		Completed: task.Completed,
	}
}
