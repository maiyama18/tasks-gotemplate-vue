package database

import (
	"fmt"

	"github.com/maiyama18/tasks-gotemplate-vue/domain/model"
)

type TaskInMemoryDatabase struct {
	nextID uint64
	tasks  []*model.Task
}

func NewTaskInMemoryDatabase() *TaskInMemoryDatabase {
	initialTasks := []*model.Task{
		{ID: 1, Title: "Learn Visa", Completed: false},
		{ID: 2, Title: "Write Vue.js App", Completed: false},
	}

	return &TaskInMemoryDatabase{
		nextID: 3,
		tasks:  initialTasks,
	}
}

func (t *TaskInMemoryDatabase) Find(id uint64) (*model.Task, error) {
	for _, t := range t.tasks {
		if t.ID == id {
			return t, nil
		}
	}
	return nil, fmt.Errorf("task with id %d not found", id)
}

func (t *TaskInMemoryDatabase) FindAll() ([]*model.Task, error) {
	return t.tasks, nil
}

func (t *TaskInMemoryDatabase) Create(task *model.Task) error {
	task.ID = t.nextID
	t.nextID++

	t.tasks = append(t.tasks, task)
	return nil
}

func (t *TaskInMemoryDatabase) Update(task *model.Task) error {
	for _, t := range t.tasks {
		if t.ID == task.ID {
			t.UpdateAttributes(task)
			return nil
		}
	}
	return fmt.Errorf("task with id %d not found", task.ID)
}
