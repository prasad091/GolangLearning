package class

import "testing"

func TestNewTask(t *testing.T) {
	task := NewTask("Learn Go")
	if task.Title != "Learn Go" {
		t.Errorf("expected learn go, got %v", task.Title)
	}
}
