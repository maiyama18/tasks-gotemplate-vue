package model

type Task struct {
	ID        uint64
	Title     string
	Completed bool
}

func (t *Task) Toggle() {
	t.Completed = !t.Completed
}
