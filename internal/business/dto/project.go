package dto

import "task-manager-go/internal/domain"

type CreateProjectRequest struct {
	Name string
}

type ProjectDTO struct {
	ID     domain.ProjectID
	Name   string
	UserID domain.UserID
}

func (p ProjectDTO) String() string {
	return "Project " + p.ID.String() + "-" + p.Name
}

func ToProjectDTO(e domain.Project) ProjectDTO {
	return ProjectDTO{
		ID:     e.ID,
		Name:   e.Name,
		UserID: e.UserId,
	}
}
