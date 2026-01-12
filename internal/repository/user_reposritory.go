package repository

import "task-manager-go/internal/domain"

type UserRepository interface {
	Save(u domain.User) error
	Get(id domain.UserID) (domain.User, error)
	FindByLogin(login string) (domain.User, error)
	All() []domain.User
}
