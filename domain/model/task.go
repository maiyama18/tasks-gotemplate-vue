package model

type Task struct {
	ID        uint64
	Title     string
	Completed bool
}

func (t *Task) Complete() {
	t.Completed = true
}
