package project

import (
	"task-manager-go/internal/business/dto"
	"task-manager-go/internal/business/service/task"
	"task-manager-go/internal/domain"
	"task-manager-go/internal/repository"
)

type ProjectService interface {
	Create(userID domain.UserID, request dto.CreateProjectRequest) (domain.ProjectID, error)
	GetById(id domain.ProjectID) (dto.ProjectDTO, error)
	FindByUserId(id domain.UserID) ([]dto.ProjectDTO, error)
	DeleteById(id domain.ProjectID) error
	AddTask(request dto.CreateTaskRequest) (domain.TaskID, error)
	GetTasks(projectID domain.ProjectID) ([]dto.TaskDTO, error)
}

type ProjectServiceImpl struct {
	projectRepository repository.ProjectRepository
	taskService       task.TaskService
}

func NewProjectService(projectRepository repository.ProjectRepository, taskService task.TaskService) *ProjectServiceImpl {
	return &ProjectServiceImpl{projectRepository: projectRepository, taskService: taskService}
}

func (s *ProjectServiceImpl) Create(userID domain.UserID, request dto.CreateProjectRequest) (domain.ProjectID, error) {
	project := domain.Project{
		ID:     domain.NewProjectId(),
		Name:   request.Name,
		UserId: userID,
	}

	err := s.projectRepository.Save(project)
	if err != nil {
		return domain.ProjectID{}, err
	}
	return project.ID, nil
}

func (s *ProjectServiceImpl) GetById(id domain.ProjectID) (dto.ProjectDTO, error) {
	project, err := s.projectRepository.Get(id)
	return dto.ToProjectDTO(project), err
}

func (s *ProjectServiceImpl) FindByUserId(userId domain.UserID) ([]dto.ProjectDTO, error) {
	projects, err := s.projectRepository.FindByUser(userId)
	if err != nil {
		return nil, err
	}
	dtos := []dto.ProjectDTO{}
	for _, project := range projects {
		dtos = append(dtos, dto.ToProjectDTO(project))
	}
	return dtos, nil
}

func (s *ProjectServiceImpl) DeleteById(id domain.ProjectID) error {
	err := s.taskService.DeleteByProjectID(id)
	if err != nil {
		return err
	}
	return s.projectRepository.DeleteById(id)
}

func (s *ProjectServiceImpl) AddTask(request dto.CreateTaskRequest) (domain.TaskID, error) {
	_, err := s.projectRepository.Get(request.ProjectID)
	if err != nil {
		return domain.TaskID{}, err
	}
	return s.taskService.Create(request)
}

func (s *ProjectServiceImpl) GetTasks(projectID domain.ProjectID) ([]dto.TaskDTO, error) {
	_, err := s.projectRepository.Get(projectID)
	if err != nil {
		return nil, err
	}
	tasks, err := s.taskService.FindByProjectID(projectID)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
