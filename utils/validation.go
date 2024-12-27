package utils

import "fmt"

//It ensures that the data provided for creating or updating tasks meets specific criteria
// ValidateTaskInput - Validate input for creating or updating a task

// validation should be handled seperately from business logic. By placing validation login in validation.go
func ValidateTaskInput(title, description string) error {
    if title == "" || description == "" {
        return fmt.Errorf("both title and description are required")
    }
    return nil
}
