// handlers.go
// This file implements HTTP handlers for managing tasks in the task manager system. 
// The handlers interact with the db package to perform CRUD operations on tasks, 
// including creating, reading, updating, deleting, and filtering tasks. 
// The code also includes CLI handlers to manage tasks outside the HTTP context.
// Each handler returns appropriate HTTP responses with status codes, 
// and handles validation errors and database interaction errors gracefully.

package handlers

import (
	"TASK-MANAGER/db"
	"TASK-MANAGER/models"
	"TASK-MANAGER/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"        // A router package to manage URL routing and extract variables like id from the path 
)

// CreateTask - Handle creating a task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task      // Initializa a variable to hold the task data. It maps diretly to the structure defined in model

	// Decode the json payload sent in the request body in the task variable  if the client sends malformed JSON,an error is introduced
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Validate input ensures that required fiels Title and description are present and valid 
	if err := utils.ValidateTaskInput(task.Title, task.Description); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save task
	err := db.CreateTask(task)
	if err != nil {
		http.Error(w, "Error creating task", http.StatusInternalServerError)
		return
	}

	// Respond with the created task
	//Sets the response header to application/json and sends back the created task as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// CreateTaskCLI - Handle creating a task (for CLI)
func CreateTaskCLI(title, description string) (*models.Task, error) {
	// Create a task pointer
	task := &models.Task{
		Title:       title,
		Description: description,
	}

	// Validate input
	if err := utils.ValidateTaskInput(task.Title, task.Description); err != nil {
		return nil, err
	}

	// Save task to DB
	err := db.CreateTask(*task) // Dereference the pointer when saving
	if err != nil {
		return nil, fmt.Errorf("error creating task: %v", err)
	}

	return task, nil // Return the pointer to task
}

// GetAllTasks - Handle getting all tasks
func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.GetAllTasks()
	if err != nil {
		http.Error(w, "Error retrieving tasks", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// GetTaskByID - Handle getting a task by ID (for HTTP)
func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	// Extract the task ID from the URL parameters
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)     //convert the id string to an integer
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task, err := db.GetTaskByID(id)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// GetTaskByIDCLI - Handle getting a task by ID (for CLI)
func GetTaskByIDCLI(id int) (*models.Task, error) {
	task, err := db.GetTaskByID(id)
	if err != nil {
		return nil, fmt.Errorf("task not found: %v", err)
	}
	return &task, nil
}

// UpdateTask - Handle updating a task
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	// Extract the task ID from the URL parameters
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	task.ID = id

	// Validate input
	if err := utils.ValidateTaskInput(task.Title, task.Description); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the task in the database
	if err := db.UpdateTask(task); err != nil {
		http.Error(w, "Error updating task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteTask - Handle deleting a task
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Extract the task ID from the URL parameters
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Delete the task from the database
	if err := db.DeleteTask(id); err != nil {
		http.Error(w, "Error deleting task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetFilteredTasks - Handle task filtering by category and status
func GetFilteredTasks(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	status := r.URL.Query().Get("status")

	tasks, err := db.GetFilteredTasks(category, status)
	if err != nil {
		http.Error(w, "Error retrieving tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
