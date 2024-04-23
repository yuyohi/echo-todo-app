package app

import "time"

// write type from todo
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewTask(title, description, status string) *Task {
	return &Task{
		Title:       title,
		Description: description,
		Status:      status,
		CreatedAt:   time.Now(),
	}
}
