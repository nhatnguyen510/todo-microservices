package internal

import "time"

type Priority string

const (
	Low    Priority = "low"
	Medium Priority = "medium"
	High   Priority = "high"
)

type Todo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Priority  Priority  `json:"priority"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
