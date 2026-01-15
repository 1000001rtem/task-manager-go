package dto

import "task-manager-go/internal/domain"

type CreateTaskRequest struct {
	ProjectID domain.ProjectID
	Name      string
}

type TaskDTO struct {
	TaskID    domain.TaskID
	Name      string
	ProjectID domain.ProjectID
	Status    domain.TaskStatus
}

func (t TaskDTO) String() string {
	return "Task " + t.TaskID.String() + "-" + t.Name
}

func ToTaskDTO(task domain.Task) TaskDTO {
	return TaskDTO{
		TaskID:    task.ID,
		Name:      task.Name,
		ProjectID: task.ProjectID,
		Status:    task.Status,
	}
}
