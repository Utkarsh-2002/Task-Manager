
// db.go
// This file implements an in-memory task manager with basic CRUD operations: 
// Create, Read, Update, Delete, and filtering tasks based on category and status. 
// Tasks are stored in a global slice `tasks` and accessed through synchronized methods 
// to ensure thread safety in a concurrent environment. The operations are guarded 
// by a mutex (`mu`) to ensure that only one goroutine can access or modify the tasks 
// slice at a time, preventing race conditions.


package db

import (
	"TASK-MANAGER/models"
	"fmt"
	"sync"
	
)


var tasks = []models.Task{}


var mu sync.Mutex // A mutex used to ensure that only one goroutine can access or modify the tasks slice at a time, preventing rare conditions in a concurrent enviromentation
func ClearTasks() {
	mu.Lock()
	defer mu.Unlock()
	tasks = []models.Task{} // Reset the tasks slice
}

// GetAllTasks - Get all tasks    Returns all tasks in the tasks slice
func GetAllTasks() ([]models.Task, error) {
	mu.Lock()
	defer mu.Unlock() // Lock the tasks slice to prevent concurrent modifications while reading
	return tasks, nil
}

// GetTaskByID - Get a task by ID
func GetTaskByID(id int) (models.Task, error) {
	mu.Lock()
	defer mu.Unlock()
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return models.Task{}, fmt.Errorf("task not found")
}

// CreateTask - Create a new task
func CreateTask(task models.Task) error {
	mu.Lock()
	defer mu.Unlock()
	task.ID = len(tasks) + 1
	tasks = append(tasks, task) // add the task to the tasks slic
	return nil
}

// UpdateTask - Update an existing task
func UpdateTask(task models.Task) error {
	mu.Lock()
	defer mu.Unlock()
	for i, t := range tasks {
		if t.ID == task.ID {
			tasks[i] = task
			return nil
		}
	}
	return fmt.Errorf("task not found")
}

// DeleteTask - Delete a task by ID
func DeleteTask(id int) error {
	mu.Lock()
	defer mu.Unlock()
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...) // all tasks before the targer and all tasks after the target
			return nil
		}
	}
	return fmt.Errorf("task not found")
}

// GetFilteredTasks - Get tasks filtered by category and status
func GetFilteredTasks(category, status string) ([]models.Task, error) {
	mu.Lock()
	defer mu.Unlock()
	var filteredTasks []models.Task
	for _, task := range tasks {
		if (category == "" || task.Category == category) && (status == "" || task.Status == status) {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks, nil
}
