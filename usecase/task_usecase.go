package usecase

import (
	"fmt"

	"github.com/maiyama18/tasks-gotemplate-vue/domain/model"
	"github.com/maiyama18/tasks-gotemplate-vue/domain/repository"
)

type TaskUsecase struct {
	taskRepository repository.TaskRepository
}

func NewTaskUsecase(taskRepository repository.TaskRepository) *TaskUsecase {
	return &TaskUsecase{taskRepository: taskRepository}
}

func (t *TaskUsecase) FindAll() ([]*model.Task, error) {
	return t.taskRepository.FindAll()
}

func (t *TaskUsecase) Add(title string) (*model.Task, error) {
	task := &model.Task{Title: title, Completed: false}
	return t.taskRepository.Create(task)
}

func (t *TaskUsecase) Complete(id uint64) (*model.Task, error) {
	task, err := t.taskRepository.Find(id)
	if err != nil {
		return nil, err
	}

	if task.Completed {
		return nil, fmt.Errorf("task with id %d is already completed", id)
	}

	task.Complete()
	return t.taskRepository.Update(task)
}
