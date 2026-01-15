package memory

import (
	"maps"
	"sync"
	"task-manager-go/internal/domain"
	"task-manager-go/internal/repository"
	"task-manager-go/internal/util"
)

type userStorage struct {
	mu      sync.RWMutex
	storage map[domain.UserID]*domain.User
}

func NewUserStorage() repository.UserRepository {
	s := &userStorage{
		storage: make(map[domain.UserID]*domain.User),
	}
	id := domain.NewUserId()
	s.storage[id] = &domain.User{
		ID:       id,
		Login:    "Test",
		Email:    "Test@Test.com",
		Password: util.HashPassword("Test"),
	}
	return s
}

func (s *userStorage) Save(user domain.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, err := s.findByLogin(user.Login)
	if err == nil {
		return domain.ErrAlreadyExists
	}
	s.storage[user.ID] = &user
	return nil
}

func (s *userStorage) Get(id domain.UserID) (domain.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	user, ok := s.storage[id]
	if !ok {
		return domain.User{}, domain.FormatError(domain.ErrEntityNotFound, "User with id: %s not found", id)
	}
	return *user, nil
}

func (s *userStorage) FindByLogin(login string) (domain.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.findByLogin(login)
}

func (s *userStorage) All() []domain.User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	users := make([]domain.User, 0, len(s.storage))
	for u := range maps.Values(s.storage) {
		users = append(users, *u)
	}
	return users
}

func (s *userStorage) findByLogin(login string) (domain.User, error) {
	for u := range maps.Values(s.storage) {
		if u.Login == login {
			return *u, nil
		}
	}
	return domain.User{}, domain.FormatError(domain.ErrEntityNotFound, "User with login: %s not found", login)
}
