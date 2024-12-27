// db_test.go
// This file contains unit tests for the CRUD operations implemented in the db package. 
// It tests the functionality of task management, including task creation, reading, updating, 
// deleting, and filtering tasks. The tests use the testify/assert package for asserting expected results.
// Each test ensures that the database operations function correctly, handling edge cases like 
// non-existing tasks, empty task lists, and task updates or deletions. The tasks slice is cleared 
// before each test to ensure a clean state for each individual test, ensuring test isolation and consistency.

package db_test

import (
	"TASK-MANAGER/db"
	"TASK-MANAGER/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

//var tt []models.Task

// TestMain will run before all tests to clear tasks slice
func TestMain(m *testing.M) {
	// Clear tasks before running tests
	db.ClearTasks()
	m.Run()
}

// Helper function to clear tasks before each test

func TestGetTaskByID(t *testing.T) {
	// Clear tasks before the test
	db.ClearTasks()

	// First, create a task to ensure there's at least one task in the list
	task := models.Task{Title: "Test Task", Description: "Test Description"}
	err := db.CreateTask(task)
	assert.NoError(t, err)

	// Now, fetch the task by its ID (which will be 1 after creation)
	task, err = db.GetTaskByID(1) // Use ID 1, which is guaranteed to exist
	assert.NoError(t, err)
	assert.Equal(t, "Test Task", task.Title)
}

func TestGetAllTasks(t *testing.T) {
	// Clear tasks before the test
	db.ClearTasks()

	// First, create a task to ensure there's at least one task in the list
	task := models.Task{Title: "Test Task", Description: "Test Description"}
	err := db.CreateTask(task)
	assert.NoError(t, err)

	// Now, fetch all tasks using the global tasks variable
	result, err := db.GetAllTasks() // Directly call the function
	assert.NoError(t, err)
	assert.Len(t, result, 1) // Ensure the slice has 1 task
}

func TestUpdateTask(t *testing.T) {
	// Clear tasks before the test
	db.ClearTasks()

	// First, create a task to ensure there's a task to update
	task := models.Task{Title: "Test Task", Description: "Test Description"}
	err := db.CreateTask(task)
	assert.NoError(t, err)

	// Now, update the task
	updatedTask := models.Task{ID: 1, Title: "Updated Task", Description: "Updated Description"}
	err = db.UpdateTask(updatedTask) // Directly call the function
	assert.NoError(t, err)

	// Verify that the task has been updated
	taskFromDB, err := db.GetTaskByID(1)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Task", taskFromDB.Title)
	assert.Equal(t, "Updated Description", taskFromDB.Description)
}

func TestDeleteTask(t *testing.T) {
	// Clear tasks before the test
	db.ClearTasks()

	// First, create a task to ensure there's a task to delete
	task := models.Task{Title: "Test Task", Description: "Test Description"}
	err := db.CreateTask(task)
	assert.NoError(t, err)

	// Now, delete the task with ID 1
	err = db.DeleteTask(1) // Directly call the function
	assert.NoError(t, err)

	// Verify that the task has been deleted by attempting to fetch it
	_, err = db.GetTaskByID(1)
	assert.Error(t, err) // Expect an error because the task should be deleted
}

func TestGetFilteredTasks(t *testing.T) {
	// Clear tasks before the test
	db.ClearTasks()

	// Add a task with category "Test" and status "Completed" first
	task := models.Task{
		Title:       "Test Task",
		Category:    "Test",
		Status:      "Completed",
		Description: "This is a test task",
	}
	err := db.CreateTask(task)
	assert.NoError(t, err)

	// Now test the filter
	result, err := db.GetFilteredTasks("Test", "Completed")
	assert.NoError(t, err)
	assert.Len(t, result, 1) // Ensure 1 task matches the filter
}

func TestGetAllTasksEmpty(t *testing.T) {
	// Clear tasks before the test
	db.ClearTasks()

	// Test that the task list is empty initially
	result, err := db.GetAllTasks()
	assert.NoError(t, err)
	assert.Len(t, result, 0) // Ensure the slice is empty
}
func TestGetTaskByIDNotFound(t *testing.T) {
	// Clear tasks before the test
	db.ClearTasks()

	// Try fetching a non-existing task
	_, err := db.GetTaskByID(999) // ID 999 doesn't exist
	assert.Error(t, err) // Expect an error because the task doesn't exist
}
func TestUpdateTaskNotFound(t *testing.T) {
	// Clear tasks before the test
	db.ClearTasks()

	// Try updating a task that doesn't exist
	updatedTask := models.Task{ID: 999, Title: "Non-existent", Description: "This task doesn't exist"}
	err := db.UpdateTask(updatedTask)
	assert.Error(t, err) // Expect an error because the task doesn't exist
}


func TestDeleteTaskNotFound(t *testing.T) {
	// Clear tasks before the test
	db.ClearTasks()

	// Try deleting a task that doesn't exist
	err := db.DeleteTask(999) // ID 999 doesn't exist
	assert.Error(t, err) // Expect an error because the task doesn't exist
}

func TestClearTasks(t *testing.T) {
	// Clear tasks before the test
	db.ClearTasks()

	// Add some tasks
	task1 := models.Task{Title: "Task 1", Description: "Description 1"}
	task2 := models.Task{Title: "Task 2", Description: "Description 2"}
	db.CreateTask(task1)
	db.CreateTask(task2)

	// Clear tasks
	db.ClearTasks()

	// Ensure tasks list is empty after clearing
	result, err := db.GetAllTasks()
	assert.NoError(t, err)
	assert.Len(t, result, 0) // Ensure the tasks slice is empty
}






