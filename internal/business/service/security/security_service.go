package security

import (
	"task-manager-go/internal/business/dto"
	"task-manager-go/internal/domain"
	"task-manager-go/internal/repository"
	"task-manager-go/internal/util"
)

type SecurityService interface {
	Login(login string, password string) error
}

type securityService struct {
	securityContext *context
	userRepository  repository.UserRepository
}

func NewSecurityService(userRepository repository.UserRepository) SecurityService {
	return &securityService{securityContext: SecurityContext, userRepository: userRepository}
}

func (s *securityService) Login(login string, password string) error {
	user, err := s.userRepository.FindByLogin(login)
	if err != nil {
		return domain.ErrUserNotFound
	}

	if util.ComparePasswords(user.Password, password) {
		s.securityContext.setCurrentUser(dto.ToUserDTO(user))
		return nil
	}
	return domain.ErrWrongPassword
}
