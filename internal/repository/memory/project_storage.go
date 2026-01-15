package memory

import (
	"sync"
	"task-manager-go/internal/domain"
	"task-manager-go/internal/repository"
)

type projectStorage struct {
	mutex   sync.RWMutex
	storage map[domain.ProjectID]*domain.Project
}

func NewProjectStorage() repository.ProjectRepository {
	return &projectStorage{
		storage: make(map[domain.ProjectID]*domain.Project),
	}
}

func (s *projectStorage) Save(project domain.Project) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.storage[project.ID] = &project
	return nil
}

func (s *projectStorage) Get(id domain.ProjectID) (domain.Project, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	project, ok := s.storage[id]
	if !ok {
		return domain.Project{}, domain.FormatError(domain.ErrEntityNotFound, "Project with id: %s not found", id)
	}
	return *project, nil
}

func (s *projectStorage) FindByUser(userID domain.UserID) ([]domain.Project, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	var projects []domain.Project
	for _, project := range s.storage {
		if project.UserId == userID {
			projects = append(projects, *project)
		}
	}
	return projects, nil
}

func (s *projectStorage) DeleteById(id domain.ProjectID) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.storage, id)
	return nil
}
