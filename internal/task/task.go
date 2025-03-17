package task

import (
	"time"
)

// Task struct represents a task in the to-do list.
type Task struct {
	ID         int       `csv:"id"`
	Name       string    `csv:"name"`
	Done       bool      `csv:"done"`
	CreatedAt  time.Time `csv:"created_at"`
	ModifiedAt time.Time `csv:"modified_at"`
}

// Filter struct represents a filter that can be used to filter tasks.
type Filter struct {
	Include string
	OrderBy string
	Desc    bool
	Asc     bool
	Done    string
	Limit   int
}

// Service interface defines methods that can be used to interact with tasks.
type Service interface {
	List(filter Filter) ([]Task, error)
	Create(name string) (Task, error)
	Update(id int, newName string) (Task, error)
	Complete(id int) (Task, error)
	Undone(id int) (Task, error)
	Delete(id int) error
}
