package models

import "time"

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Category    string    `json:"category"`
	DueDate     time.Time `json:"due_date"`
}
