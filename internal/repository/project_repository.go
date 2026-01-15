package repository

import "task-manager-go/internal/domain"

type ProjectRepository interface {
	Save(domain.Project) error
	Get(domain.ProjectID) (domain.Project, error)
	FindByUser(domain.UserID) ([]domain.Project, error)
	DeleteById(domain.ProjectID) error
}
