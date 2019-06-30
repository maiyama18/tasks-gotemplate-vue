package repository

import "github.com/maiyama18/tasks-gotemplate-vue/domain/model"

type TaskRepository interface {
	Find(id uint64) (*model.Task, error)
	FindAll() ([]*model.Task, error)
	Create(task *model.Task) (*model.Task, error)
	Update(task *model.Task) (*model.Task, error)
}
