package task

// Store defines how tasks are persisted and retrieved
type Store interface {
	// GetAll retrieves all tasks from storage
	GetAll(filter *Filter) ([]Task, error)

	// Get retrieves a single task by ID
	Get(id int) (*Task, error)

	// Save persists a new task
	Save(task *Task) error

	// SaveAll persists multiple tasks
	SaveAll(tasks []Task) error

	// Update updates an existing task
	Update(id int, task *Task) error

	// Delete removes a task from storage
	Delete(id int) error

	// NextID generates the next available ID for a new task
	NextID() (int, error)
}
