package domain

import "github.com/google/uuid"

type ProjectID uuid.UUID

func (id ProjectID) String() string {
	return uuid.UUID(id).String()
}

func NewProjectId() ProjectID {
	return ProjectID(uuid.New())
}

type Project struct {
	ID     ProjectID
	Name   string
	UserId UserID
}
