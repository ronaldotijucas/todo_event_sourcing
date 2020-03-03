package task

import "github.com/ronaldotijucas/todo/es"

type Task struct {
}

func (t Task) Create(command *CreateTaskCmd) ([]*es.Event, error) {
	return []*es.Event{
		{
			AggregateID: command.ID,
			Payload: &TaskCreated{
				ID:          command.ID,
				Description: command.Description,
				CreatedAt:   command.CreateAt,
			},
		},
	}, nil
}

