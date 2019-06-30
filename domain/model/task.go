package model

type Task struct {
	ID        uint64
	Title     string
	Completed bool
}

func (t *Task) Complete() {
	t.Completed = true
}

func (t *Task) UpdateAttributes(newTask *Task) {
	t.Title = newTask.Title
	t.Completed = newTask.Completed
}
