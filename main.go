package main

import (
	"TASK-MANAGER/cli"
	"TASK-MANAGER/handlers" // Corrected import
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Run CLI commands first if needed
	cli.Execute()

	// Create the router using gorilla/mux
	router := mux.NewRouter()             // the router is responsible for matching incoming HTTP requests to their respective handler functions.

	// API Routes
	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")             // Create a task
	router.HandleFunc("/tasks", handlers.GetAllTasks).Methods("GET")             // Get all tasks
	router.HandleFunc("/tasks/{id}", handlers.GetTaskByID).Methods("GET")        // Get a task by ID
	router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")         // Update a task by ID
	router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")      // Delete a task by ID
	router.HandleFunc("/tasks/filter", handlers.GetFilteredTasks).Methods("GET") // Get filtered tasks

	// Start the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}
