package user

import (
	"task-manager-go/internal/business/dto"
	"task-manager-go/internal/domain"
	"task-manager-go/internal/repository"
	"task-manager-go/internal/util"
)

type UserService interface {
	Create(request dto.CreateUserRequest) (domain.UserID, error)
	GetById(id domain.UserID) (dto.UserDTO, error)
	FindAll() []dto.UserDTO
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) Create(request dto.CreateUserRequest) (domain.UserID, error) {
	if util.IsEmpty(request.Email) || util.IsEmpty(request.Login) {
		return domain.UserID{}, domain.FormatError(domain.ErrIllegalArgument, "Input arguments cant be empty")
	}
	user := domain.User{
		ID:       domain.NewUserId(),
		Login:    request.Login,
		Email:    request.Email,
		Password: util.HashPassword(request.Password),
	}
	if err := s.userRepository.Save(user); err != nil {
		return domain.UserID{}, err
	}
	return user.ID, nil
}

func (s *userService) GetById(id domain.UserID) (dto.UserDTO, error) {
	u, err := s.userRepository.Get(id)
	return dto.ToUserDTO(u), err
}

func (s *userService) FindAll() []dto.UserDTO {
	users := s.userRepository.All()
	dtos := []dto.UserDTO{}
	for _, u := range users {
		dtos = append(dtos, dto.ToUserDTO(u))
	}
	return dtos
}
