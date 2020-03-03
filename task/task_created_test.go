package task_test

import (
	"github.com/ronaldotijucas/todo/task"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)


func Test_it_creates_a_new_task(t *testing.T) {
	model := &task.Task{}
	events, err := model.Create(&task.CreateTaskCmd{
		ID: "uuid",
		Description: "Fix breakfast",
		CreateAt: time.Date(2020,2, 28, 0, 0, 0,0, time.UTC),
	})
	assert.NoError(t, err)

	actual:= events[0].Payload.(*task.TaskCreated)
	assert.Equal(t, 1, len(events))
	assert.Equal(t, &task.TaskCreated{
		ID: "uuid",
		Description: "Fix breakfast",
		CreatedAt: time.Date(2020,2,28,0,0,0,0, time.UTC),
	}, actual)
}