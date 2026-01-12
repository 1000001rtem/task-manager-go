package repository

import "task-manager-go/internal/domain"

type TaskRepository interface {
	Save(domain.Task) error
	Get(domain.TaskID) (domain.Task, error)
	FindByProject(domain.ProjectID) []domain.Task
	DeleteById(domain.TaskID) error
	Update(domain.Task) (domain.Task, error)
}
