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

func (t *TaskUsecase) Add(task *model.Task) error {
	return t.taskRepository.Create(task)
}

func (t *TaskUsecase) Complete(id uint64) error {
	task, err := t.taskRepository.Find(id)
	if err != nil {
		return err
	}

	if task.Completed {
		return fmt.Errorf("task with id %d is already completed", id)
	}

	task.Complete()
	return t.taskRepository.Update(task)
}
