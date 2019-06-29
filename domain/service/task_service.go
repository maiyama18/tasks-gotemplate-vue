package service

import "github.com/maiyama18/tasks-gotemplate-vue/domain/model"

type TaskService interface {
	FindAll() ([]*model.Task, error)
	Add(task *model.Task) error
	Complete(id uint64) error
}
