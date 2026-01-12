package memory

import "task-manager-go/internal/repository"

type Store struct {
	UserRepository    repository.UserRepository
	ProjectRepository repository.ProjectRepository
	TaskRepository    repository.TaskRepository
}

func NewStore() *Store {
	return &Store{
		UserRepository:    NewUserStorage(),
		ProjectRepository: NewProjectStorage(),
		TaskRepository:    NewTaskStorage(),
	}
}
