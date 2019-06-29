package repository

import "github.com/maiyama18/tasks-gotemplate-vue/domain/model"

type TaskRepository interface {
	FindAll() ([]model.Task, error)
	Create(task model.Task) error
	Update(task model.Task) error
}
