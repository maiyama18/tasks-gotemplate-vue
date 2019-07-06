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

func (t *TaskInMemoryDatabase) Create(task *model.Task) (*model.Task, error) {
	task.ID = t.nextID
	t.nextID++

	t.tasks = append(t.tasks, task)
	return task, nil
}

func (t *TaskInMemoryDatabase) Update(task *model.Task) (*model.Task, error) {
	for _, t := range t.tasks {
		if t.ID == task.ID {
			t.Title = task.Title
			t.Completed = task.Completed
			return t, nil
		}
	}
	return nil, fmt.Errorf("task with id %d not found", task.ID)
}

func (t *TaskInMemoryDatabase) Delete(id uint64) error {
	_, err := t.Find(id)
	if err != nil {
		return err
	}

	var updatedTasks []*model.Task
	for _, t := range t.tasks {
		if t.ID == id {
			continue
		}

		updatedTasks = append(updatedTasks, t)
	}
	t.tasks = updatedTasks

	return nil
}
