package class

type Task struct {
	Title string
}

func NewTask(title string) *Task {
	return &Task{title}
}
