package domain

import "github.com/google/uuid"

type TaskID uuid.UUID

type TaskStatus string

func (id TaskID) String() string {
	return uuid.UUID(id).String()
}

func NewTaskId() TaskID {
	return TaskID(uuid.New())
}

const (
	Todo       TaskStatus = "TODO"
	InProgress TaskStatus = "IN_PROGRESS"
	Done       TaskStatus = "DONE"
)

type Task struct {
	ID        TaskID
	ProjectID ProjectID
	Name      string
	Status    TaskStatus
}
