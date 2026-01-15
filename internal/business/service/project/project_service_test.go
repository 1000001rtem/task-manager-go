package project

import (
	"errors"
	"task-manager-go/internal/business/dto"
	"task-manager-go/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProjectServiceImpl_AddTask(t *testing.T) {
	tests := []struct {
		name                       string
		request                    dto.CreateTaskRequest
		gotErr                     error
		mockProjectRepositorySetup func(*MockProjectRepository)
		mockTaskServiceSetup       func(*MockTaskService)
	}{
		{
			"should success create task",
			dto.CreateTaskRequest{Name: "test", ProjectID: domain.NewProjectId()},
			nil,
			func(r *MockProjectRepository) { r.On("Get", mock.Anything).Return(domain.Project{}, nil).Once() },
			func(r *MockTaskService) { r.On("Save", mock.Anything).Return(domain.NewTaskId(), nil).Once() },
		},
		{
			"should return error when project not exist",
			dto.CreateTaskRequest{Name: "test", ProjectID: domain.NewProjectId()},
			domain.ErrEntityNotFound,
			func(r *MockProjectRepository) {
				r.On("Get", mock.Anything).Return(domain.Project{}, domain.ErrEntityNotFound).Once()
			},
			func(r *MockTaskService) {},
		},
		{
			"should return error when task project return error",
			dto.CreateTaskRequest{Name: "test", ProjectID: domain.NewProjectId()},
			errors.New("error"),
			func(r *MockProjectRepository) { r.On("Get", mock.Anything).Return(domain.Project{}, nil).Once() },
			func(r *MockTaskService) {
				r.On("Save", mock.Anything).Return(domain.TaskID{}, errors.New("error")).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProjectRepository := new(MockProjectRepository)
			mockTaskService := new(MockTaskService)
			tt.mockProjectRepositorySetup(mockProjectRepository)
			tt.mockTaskServiceSetup(mockTaskService)
			s := &ProjectServiceImpl{
				projectRepository: mockProjectRepository,
				taskService:       mockTaskService,
			}
			result, err := s.AddTask(tt.request)
			if tt.gotErr != nil {
				assert.ErrorAs(t, err, &tt.gotErr)
			} else {
				assert.NotNil(t, result)
			}
			mockProjectRepository.AssertExpectations(t)
			mockTaskService.AssertExpectations(t)
		})
	}
}

func TestProjectServiceImpl_Create(t *testing.T) {
	tests := []struct {
		name      string
		request   dto.CreateProjectRequest
		gotErr    error
		mockSetup func(*MockProjectRepository)
	}{
		{
			"should success create project",
			dto.CreateProjectRequest{Name: "test"},
			nil,
			func(r *MockProjectRepository) { r.On("Save", mock.Anything).Return(nil).Once() },
		},
		{
			"should throw error when repository fails",
			dto.CreateProjectRequest{Name: "test"},
			errors.New("test"),
			func(r *MockProjectRepository) { r.On("Save", mock.Anything).Return(errors.New("test")).Once() },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProjectRepository := new(MockProjectRepository)
			mockTaskService := new(MockTaskService)
			tt.mockSetup(mockProjectRepository)
			s := &ProjectServiceImpl{
				projectRepository: mockProjectRepository,
				taskService:       mockTaskService,
			}
			result, err := s.Create(domain.NewUserId(), tt.request)
			if tt.gotErr != nil {
				assert.ErrorAs(t, err, &tt.gotErr)
			} else {
				assert.NotNil(t, result)
			}
			mockProjectRepository.AssertExpectations(t)
		})
	}
}

func TestProjectServiceImpl_DeleteById(t *testing.T) {
	id := domain.NewProjectId()
	tests := []struct {
		name                       string
		request                    domain.ProjectID
		gotErr                     error
		mockProjectRepositorySetup func(*MockProjectRepository)
		mockTaskServiceSetup       func(*MockTaskService)
	}{
		{
			"should delete with tasks",
			id,
			nil,
			func(r *MockProjectRepository) { r.On("DeleteById", id).Return(nil).Once() },
			func(r *MockTaskService) { r.On("DeleteByProjectID", id).Return(nil).Once() },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProjectRepository := new(MockProjectRepository)
			mockTaskService := new(MockTaskService)
			tt.mockProjectRepositorySetup(mockProjectRepository)
			tt.mockTaskServiceSetup(mockTaskService)
			s := &ProjectServiceImpl{
				projectRepository: mockProjectRepository,
				taskService:       mockTaskService,
			}
			err := s.DeleteById(tt.request)
			if tt.gotErr != nil {
				assert.ErrorAs(t, err, &tt.gotErr)
			}
			mockProjectRepository.AssertExpectations(t)
			mockTaskService.AssertExpectations(t)
		})
	}

}

func TestProjectServiceImpl_FindByUserId(t *testing.T) {
	mockReturns := []domain.Project{{Name: "test"}, {Name: "test2"}}
	want := []dto.ProjectDTO{{Name: "test"}, {Name: "test2"}}
	tests := []struct {
		name      string
		gotErr    error
		want      []dto.ProjectDTO
		mockSetup func(*MockProjectRepository)
	}{
		{
			"should success return projects",
			nil,
			want,
			func(r *MockProjectRepository) { r.On("FindByUser", mock.Anything).Return(mockReturns, nil).Once() },
		},
		{
			"should throw error when repository fails",
			errors.New("test"),
			nil,
			func(r *MockProjectRepository) {
				r.On("FindByUser", mock.Anything).Return([]domain.Project{}, errors.New("test")).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProjectRepository := new(MockProjectRepository)
			mockTaskService := new(MockTaskService)
			tt.mockSetup(mockProjectRepository)
			s := &ProjectServiceImpl{
				projectRepository: mockProjectRepository,
				taskService:       mockTaskService,
			}
			result, err := s.FindByUserId(domain.NewUserId())
			if tt.gotErr != nil {
				assert.ErrorAs(t, err, &tt.gotErr)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, 2, len(result))
			}
			mockProjectRepository.AssertExpectations(t)
		})
	}
}

func TestProjectServiceImpl_GetById(t *testing.T) {
	id := domain.NewProjectId()
	tests := []struct {
		name      string
		gotErr    error
		want      dto.ProjectDTO
		mockSetup func(*MockProjectRepository)
	}{
		{
			"should success return project",
			nil,
			dto.ProjectDTO{ID: id, Name: "test"},
			func(r *MockProjectRepository) {
				r.On("Get", mock.Anything).Return(domain.Project{ID: id, Name: "test"}, nil).Once()
			},
		},
		{
			"should throw error when repository fails",
			errors.New("test"),
			dto.ProjectDTO{},
			func(r *MockProjectRepository) {
				r.On("Get", mock.Anything).Return(domain.Project{}, errors.New("test")).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProjectRepository := new(MockProjectRepository)
			mockTaskService := new(MockTaskService)
			tt.mockSetup(mockProjectRepository)
			s := &ProjectServiceImpl{
				projectRepository: mockProjectRepository,
				taskService:       mockTaskService,
			}
			result, err := s.GetById(domain.NewProjectId())
			if tt.gotErr != nil {
				assert.ErrorAs(t, err, &tt.gotErr)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, tt.want.ID, result.ID)
				assert.Equal(t, tt.want.Name, result.Name)
			}
			mockProjectRepository.AssertExpectations(t)
		})
	}
}

func TestProjectServiceImpl_GetTasks(t *testing.T) {
	id := domain.NewProjectId()
	want := []dto.TaskDTO{{TaskID: domain.NewTaskId(), ProjectID: id}, {TaskID: domain.NewTaskId(), ProjectID: id}}
	tests := []struct {
		name                       string
		request                    domain.ProjectID
		want                       []dto.TaskDTO
		gotErr                     error
		mockProjectRepositorySetup func(*MockProjectRepository)
		mockTaskServiceSetup       func(*MockTaskService)
	}{
		{
			"should delete return tasks",
			id,
			want,
			nil,
			func(r *MockProjectRepository) { r.On("Get", id).Return(domain.Project{ID: id}, nil).Once() },
			func(r *MockTaskService) { r.On("FindByProjectID", id).Return(want, nil).Once() },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProjectRepository := new(MockProjectRepository)
			mockTaskService := new(MockTaskService)
			tt.mockProjectRepositorySetup(mockProjectRepository)
			tt.mockTaskServiceSetup(mockTaskService)
			s := &ProjectServiceImpl{
				projectRepository: mockProjectRepository,
				taskService:       mockTaskService,
			}
			result, err := s.GetTasks(tt.request)
			if tt.gotErr != nil {
				assert.ErrorAs(t, err, &tt.gotErr)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, len(tt.want), len(result))
				assert.Equal(t, tt.want[0].TaskID, result[0].TaskID)
				assert.Equal(t, tt.want[0].ProjectID, result[0].ProjectID)
				assert.Equal(t, tt.want[1].TaskID, result[1].TaskID)
				assert.Equal(t, tt.want[1].ProjectID, result[1].ProjectID)
			}

			mockProjectRepository.AssertExpectations(t)
			mockTaskService.AssertExpectations(t)
		})
	}

}

type MockProjectRepository struct {
	mock.Mock
}

func (m *MockProjectRepository) Save(project domain.Project) error {
	err := mockMethodReturnsError("Save", &m.Mock, project)
	return err
}

func (m *MockProjectRepository) Get(id domain.ProjectID) (domain.Project, error) {
	r, e := mockMethodReturnsValueAndError("Get", &m.Mock, id)
	return r.(domain.Project), e
}

func (m *MockProjectRepository) FindByUser(id domain.UserID) ([]domain.Project, error) {
	r, e := mockMethodReturnsValueAndError("FindByUser", &m.Mock, id)
	return r.([]domain.Project), e
}

func (m *MockProjectRepository) DeleteById(id domain.ProjectID) error {
	return mockMethodReturnsError("DeleteById", &m.Mock, id)
}

type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) ChangeStatus(id domain.TaskID, status domain.TaskStatus) error {
	panic("implement me")
}

func (m *MockTaskService) Create(request dto.CreateTaskRequest) (domain.TaskID, error) {
	r, e := mockMethodReturnsValueAndError("Save", &m.Mock, request)
	return r.(domain.TaskID), e
}

func (m *MockTaskService) GetByID(id domain.TaskID) (dto.TaskDTO, error) {
	r, e := mockMethodReturnsValueAndError("GetByID", &m.Mock, id)
	return r.(dto.TaskDTO), e
}

func (m *MockTaskService) FindByProjectID(projectId domain.ProjectID) ([]dto.TaskDTO, error) {
	r, e := mockMethodReturnsValueAndError("FindByProjectID", &m.Mock, projectId)
	return r.([]dto.TaskDTO), e
}

func (m *MockTaskService) DeleteByProjectID(projectId domain.ProjectID) error {
	return mockMethodReturnsError("DeleteByProjectID", &m.Mock, projectId)
}

func mockMethodReturnsError[T any](name string, m *mock.Mock, arg T) error {
	args := m.MethodCalled(name, arg)
	return args.Error(0)
}

func mockMethodReturnsValueAndError(name string, m *mock.Mock, arg any) (any, error) {
	args := m.MethodCalled(name, arg)
	return args.Get(0), args.Error(1)
}
