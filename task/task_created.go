package task

import "time"

type TaskCreated struct {
	ID          string
	Description string
	CreatedAt   time.Time
}

