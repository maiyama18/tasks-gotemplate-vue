package router

import "github.com/maiyama18/tasks-gotemplate-vue/domain/model"

type ErrorResponse struct {
	Message string `json:"message"`
}

type TaskResponse struct {
	Task model.Task `json:"task"`
}
