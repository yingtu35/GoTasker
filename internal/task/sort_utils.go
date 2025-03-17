package task

import (
	"slices"
	"strings"
)

func SortTasks(tasks []Task, orderBy string, desc bool) []Task {
	switch orderBy {
	case "name":
		sortTasksByName(&tasks, desc)
	case "createdAt":
		sortTasksByCreatedAt(&tasks, desc)
	default:
		sortTasksByID(&tasks, desc)
	}
	return tasks
}

func sortTasksByName(tasks *[]Task, desc bool) {
	slices.SortStableFunc(*tasks, func(a, b Task) int {
		if desc {
			return strings.Compare(b.Name, a.Name)
		}
		return strings.Compare(a.Name, b.Name)
	})
}

func sortTasksByCreatedAt(tasks *[]Task, desc bool) {
	slices.SortStableFunc(*tasks, func(a, b Task) int {
		if desc {
			if b.CreatedAt.After(a.CreatedAt) {
				return 1
			} else if b.CreatedAt.Before(a.CreatedAt) {
				return -1
			}
			return 0
		}
		if a.CreatedAt.After(b.CreatedAt) {
			return 1
		} else if a.CreatedAt.Before(b.CreatedAt) {
			return -1
		}
		return 0
	})
}

func sortTasksByID(tasks *[]Task, desc bool) {
	slices.SortStableFunc(*tasks, func(a, b Task) int {
		if desc {
			return b.ID - a.ID
		}
		return a.ID - b.ID
	})
}
