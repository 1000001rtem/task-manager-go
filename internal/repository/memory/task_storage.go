package memory

import (
	"sync"
	"task-manager-go/internal/domain"
	"task-manager-go/internal/repository"
)

type taskStorage struct {
	mu      sync.RWMutex
	storage map[domain.TaskID]*domain.Task
}

func NewTaskStorage() repository.TaskRepository {
	return &taskStorage{
		storage: make(map[domain.TaskID]*domain.Task),
	}
}

func (t *taskStorage) Save(task domain.Task) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.storage[task.ID] = &task
	return nil
}

func (t *taskStorage) Get(id domain.TaskID) (domain.Task, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	task, ok := t.storage[id]
	if !ok {
		return domain.Task{}, domain.FormatError(domain.ErrEntityNotFound, "Task with id: %s not found", id)
	}
	return *task, nil
}

func (t *taskStorage) FindByProject(projectID domain.ProjectID) []domain.Task {
	t.mu.RLock()
	defer t.mu.RUnlock()
	var tasks []domain.Task
	for _, task := range t.storage {
		if task.ProjectID == projectID {
			tasks = append(tasks, *task)
		}
	}
	return tasks
}

func (t *taskStorage) DeleteById(id domain.TaskID) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	delete(t.storage, id)
	return nil
}

func (t *taskStorage) Update(task domain.Task) (domain.Task, error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.storage[task.ID] == nil {
		return domain.Task{}, domain.FormatError(domain.ErrEntityNotFound, "Task with id: %s not found", task.ID)
	}
	t.storage[task.ID] = &task
	return *t.storage[task.ID], nil
}
