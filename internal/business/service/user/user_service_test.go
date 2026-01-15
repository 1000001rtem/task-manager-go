package user

import (
	"errors"
	"task-manager-go/internal/business/dto"
	"task-manager-go/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_Create(t *testing.T) {
	tests := []struct {
		name      string
		request   dto.CreateUserRequest
		gotErr    error
		mockSetup func(*MockUserRepository)
	}{
		{
			"should create user",
			dto.CreateUserRequest{Login: "test", Email: "test"},
			nil,
			func(userRepository *MockUserRepository) {
				userRepository.On("Save", mock.Anything).Return(nil).Once()
			},
		},
		{
			"should throw error when repository fails",
			dto.CreateUserRequest{Login: "test", Email: "test"},
			errors.New("test"),
			func(userRepository *MockUserRepository) {
				userRepository.On("Save", mock.Anything).Return(errors.New("test")).Once()
			},
		},
		{
			"should throw error when login is empty",
			dto.CreateUserRequest{Login: "", Email: "test"},
			domain.ErrIllegalArgument,
			func(userRepository *MockUserRepository) {},
		},
		{
			"should throw error when email is empty",
			dto.CreateUserRequest{Login: "test", Email: ""},
			domain.ErrIllegalArgument,
			func(userRepository *MockUserRepository) {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepository := new(MockUserRepository)
			tt.mockSetup(userRepository)
			s := &userService{
				userRepository: userRepository,
			}
			result, err := s.Create(tt.request)
			if tt.gotErr != nil {
				assert.ErrorAs(t, err, &tt.gotErr)
			} else {
				assert.NotNil(t, result)
			}
			userRepository.AssertExpectations(t)
		})
	}
}

func TestUserService_GetById(t *testing.T) {
	userId := domain.NewUserId()
	tests := []struct {
		name      string
		request   domain.UserID
		want      domain.UserID
		gotErr    error
		mockSetup func(*MockUserRepository)
	}{
		{
			"should return user",
			userId,
			userId,
			nil,
			func(userRepository *MockUserRepository) {
				userRepository.On("Get", mock.Anything).Return(domain.User{ID: userId}, nil).Once()
			},
		},
		{
			"should throw error when repository fails",
			userId,
			domain.UserID{},
			errors.New("test"),
			func(userRepository *MockUserRepository) {
				userRepository.On("Get", mock.Anything).Return(domain.User{}, errors.New("test")).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepository := new(MockUserRepository)
			tt.mockSetup(userRepository)
			s := &userService{
				userRepository: userRepository,
			}
			result, err := s.GetById(tt.request)
			if tt.gotErr != nil {
				assert.ErrorAs(t, err, &tt.gotErr)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, tt.want, result.ID)
			}
			userRepository.AssertExpectations(t)
		})
	}
}

func TestUserService_FindAll(t *testing.T) {
	entities := []domain.User{
		{ID: domain.NewUserId()},
		{ID: domain.NewUserId()},
	}
	users := []dto.UserDTO{
		{ID: entities[0].ID},
		{ID: entities[1].ID},
	}
	tests := []struct {
		name      string
		want      []dto.UserDTO
		gotErr    error
		mockSetup func(*MockUserRepository)
	}{
		{
			"should return all users",
			users,
			nil,
			func(userRepository *MockUserRepository) {
				userRepository.On("All", mock.Anything).Return(entities).Once()
			},
		},
		{
			"should return empty users",
			[]dto.UserDTO{},
			errors.New("test"),
			func(userRepository *MockUserRepository) {
				userRepository.On("All", mock.Anything).Return([]domain.User(nil)).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepository := new(MockUserRepository)
			tt.mockSetup(userRepository)
			s := &userService{
				userRepository: userRepository,
			}
			result := s.FindAll()
			assert.NotNil(t, result)
			assert.Equal(t, tt.want, result)
			userRepository.AssertExpectations(t)
		})
	}
}

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Save(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Get(userID domain.UserID) (domain.User, error) {
	args := m.Called(userID)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserRepository) FindByLogin(login string) (domain.User, error) {
	args := m.Called(login)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserRepository) All() []domain.User {
	args := m.Called()
	return args.Get(0).([]domain.User)
}
