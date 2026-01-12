package security

import (
	"task-manager-go/internal/business/dto"
	"task-manager-go/internal/domain"
)

type context struct {
	currentUser *dto.UserDTO
}

var SecurityContext = &context{
	currentUser: nil,
}

func (s *context) GetCurrentUser() (dto.UserDTO, error) {
	if s.currentUser == nil {
		return dto.UserDTO{}, domain.ErrUnauthenticated
	}
	return *s.currentUser, nil
}

func (s *context) setCurrentUser(user dto.UserDTO) {
	s.currentUser = &user
}
