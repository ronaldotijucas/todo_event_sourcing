package task

import "time"

type CreateTaskCmd struct {
	ID          string
	Description string
	CreateAt    time.Time
}

