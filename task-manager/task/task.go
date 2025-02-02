package task

import "fmt"

type Task struct {
	ID          int
	Description string
	Done        bool
}

var Tasks = make([]Task, 0)

type TaskRepository interface {
	Create(task *Task) error
	Update(task *Task) error
	Delete(task *Task) error
	GetByID(id int) (*Task, error)
	GetAll() ([]Task, error)
}

func GetAll() ([]Task, error) {
	return Tasks, nil
}

func GetByID(id int) (*Task, error) {
	for _, t := range Tasks {
		if t.ID == id {
			return &t, nil
		}
	}
	return nil, fmt.Errorf("task not found")
}

func Create(description string, done bool) (*Task, error) {
	// generate a unique ID
	id := len(Tasks) + 1
	// loop through Tasks and find if it exists, if yes, update it adding a new number
	for _, t := range Tasks {
		if t.ID == id {
			id++
		}
	}
	task := Task{ID: id, Description: description, Done: done}
	Tasks = append(Tasks, task)
	return &task, nil
}

func Update(description string, done bool, id int) (*Task, error) {
	for i, t := range Tasks {
		if t.ID == id {
			Tasks[i].Description = description
			Tasks[i].Done = done
			return &Tasks[i], nil
		}
	}
	return nil, fmt.Errorf("task not found")
}

func Delete(id int) (string, error) {
	for i, t := range Tasks {
		if t.ID == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			return "task deleted", nil
		}
	}
	return "", fmt.Errorf("task not found")
}
