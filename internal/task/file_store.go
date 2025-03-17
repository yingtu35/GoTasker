package task

import (
	"log"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

// FileStore implements Store interface with CSV file
type FileStore struct {
	filePath string
}

func NewFileStore(filePath string) (Store, error) {
	taskFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Printf("could not open file: %v", err)
		return nil, err
	}
	defer taskFile.Close()

	// Check existence of file
	fileInfo, err := taskFile.Stat()
	if err != nil {
		log.Printf("could not get file info: %v", err)
		return nil, err
	}

	// If file is empty, write headers to it
	if fileInfo.Size() == 0 {
		headerContent := "id,name,done,created_at,modified_at\n"
		if _, err := taskFile.WriteString(headerContent); err != nil {
			log.Printf("could not write headers to file: %v", err)
			return nil, err
		}
	}
	return &FileStore{filePath: filePath}, nil
}

func (s *FileStore) GetAll(filter *Filter) ([]Task, error) {
	taskFile, err := os.OpenFile(s.filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Printf("could not open file: %v", err)
		return nil, err
	}
	defer taskFile.Close()

	tasks := []Task{}

	if err := gocsv.UnmarshalFile(taskFile, &tasks); err != nil {
		log.Printf("could not unmarshal file: %v", err)
		return nil, err
	}

	if filter == nil {
		return tasks, nil
	}

	filteredTasks := []Task{}
	for _, t := range tasks {
		var toInclude bool = true
		if filter.Include != "" && !strings.Contains(t.Name, filter.Include) {
			toInclude = false
		}

		if filter.Done != "" {
			if filter.Done == "true" && !t.Done {
				toInclude = false
			} else if filter.Done == "false" && t.Done {
				toInclude = false
			}
		}

		if toInclude {
			filteredTasks = append(filteredTasks, t)
		}

		SortTasks(filteredTasks, filter.OrderBy, filter.Desc)
	}

	return filteredTasks, nil
}

func (s *FileStore) Get(id int) (*Task, error) {
	tasks, err := s.GetAll(nil)
	if err != nil {
		return nil, err
	}

	for _, task := range tasks {
		if task.ID == id {
			return &task, nil
		}
	}

	return nil, nil
}

func (s *FileStore) Save(task *Task) error {
	taskFile, err := os.OpenFile(s.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer taskFile.Close()

	if err := gocsv.MarshalWithoutHeaders([]*Task{task}, taskFile); err != nil {
		return err
	}

	return nil
}

func (s *FileStore) SaveAll(tasks []Task) error {
	taskFile, err := os.OpenFile(s.filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Printf("could not open file: %v", err)
		return err
	}

	if err := gocsv.MarshalFile(&tasks, taskFile); err != nil {
		log.Printf("could not marshal file: %v", err)
		return err
	}

	return nil
}

func (s *FileStore) Update(id int, task *Task) error {
	tasks, err := s.GetAll(nil)
	if err != nil {
		return err
	}

	updatedTasks := []Task{}
	for _, t := range tasks {
		if t.ID != id {
			updatedTasks = append(updatedTasks, t)
		} else {
			updatedTasks = append(updatedTasks, *task)
		}
	}
	if err := s.SaveAll(updatedTasks); err != nil {
		return err
	}
	return nil
}

func (s *FileStore) Delete(id int) error {

	tasks, err := s.GetAll(nil)
	if err != nil {
		return err
	}

	updatedTasks := []Task{}
	for _, t := range tasks {
		if t.ID != id {
			log.Printf("task id: %d", t.ID)
			updatedTasks = append(updatedTasks, t)
		}
	}
	if err := s.SaveAll(updatedTasks); err != nil {
		return err
	}
	return nil
}

// NextID returns the next available ID for a new task
func (s *FileStore) NextID() (int, error) {
	tasks, err := s.GetAll(nil)
	if err != nil {
		return 0, err
	}

	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	return maxID + 1, nil
}
