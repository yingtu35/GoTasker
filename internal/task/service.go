package task

import (
	"log"
	"time"
)

// TaskService struct represents a service that can be used to interact with tasks.
type TaskService struct {
	store Store
}

// NewService creates a new TaskService with the provided file path.
func NewService() (Service, error) {
	store, err := NewFileStore("tasks.csv")
	if err != nil {
		log.Printf("could not create file store: %v", err)
		return nil, err
	}
	return &TaskService{store: store}, nil
}

// List retrieves tasks from the store based on the provided filter.
func (s *TaskService) List(filter Filter) ([]Task, error) {
	return s.store.GetAll(&filter)
}

func (s *TaskService) Create(name string) (Task, error) {
	nextID, err := s.store.NextID()
	if err != nil {
		return Task{}, err
	}

	newTask := Task{
		ID:         nextID,
		Name:       name,
		Done:       false,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}
	if err := s.store.Save(&newTask); err != nil {
		return Task{}, err
	}
	return newTask, nil
}

func (s *TaskService) Update(id int, newName string) (Task, error) {
	// Get task by ID
	task, err := s.store.Get(id)
	if err != nil {
		return Task{}, err
	}

	task.Name = newName
	task.ModifiedAt = time.Now()
	if err := s.store.Update(id, task); err != nil {
		return Task{}, err
	}
	return *task, nil
}

func (s *TaskService) Complete(id int) (Task, error) {
	// Get task by ID
	task, err := s.store.Get(id)
	if err != nil {
		return Task{}, err
	}
	// Mark task as done
	task.Done = true
	// Save task
	if err := s.store.Update(id, task); err != nil {
		return Task{}, err
	}

	return *task, nil
}

func (s *TaskService) Undone(id int) (Task, error) {
	// Get task by ID
	task, err := s.store.Get(id)
	if err != nil {
		return Task{}, err
	}
	// Mark task as not done
	task.Done = false
	// Save task
	if err := s.store.Update(id, task); err != nil {
		return Task{}, err
	}

	return *task, nil
}

func (s *TaskService) Delete(id int) error {
	return s.store.Delete(id)
}
