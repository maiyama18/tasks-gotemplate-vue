package usecase

import (
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

func (t *TaskUsecase) Toggle(id uint64) (*model.Task, error) {
	task, err := t.taskRepository.Find(id)
	if err != nil {
		return nil, err
	}

	task.Toggle()
	return t.taskRepository.Update(task)
}

func (t *TaskUsecase) Delete(id uint64) error {
	return t.taskRepository.Delete(id)
}
