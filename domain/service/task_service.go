package service

import "github.com/maiyama18/tasks-gotemplate-vue/domain/model"

type TaskService interface {
	FindAll() ([]*model.Task, error)
	Add(title string) (*model.Task, error)
	Toggle(id uint64) (*model.Task, error)
	Delete(id uint64) error
}
