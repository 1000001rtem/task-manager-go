package task

import (
	"errors"
	"task-manager-go/internal/business/dto"
	"task-manager-go/internal/domain"
	"task-manager-go/internal/repository"
)

type TaskService interface {
	Create(request dto.CreateTaskRequest) (domain.TaskID, error)
	GetByID(id domain.TaskID) (dto.TaskDTO, error)
	FindByProjectID(projectId domain.ProjectID) ([]dto.TaskDTO, error)
	DeleteByProjectID(projectId domain.ProjectID) error
	ChangeStatus(id domain.TaskID, status domain.TaskStatus) error
}

type TaskServiceImpl struct {
	taskRepository repository.TaskRepository
}

func NewTaskService(taskRepository repository.TaskRepository) *TaskServiceImpl {
	return &TaskServiceImpl{
		taskRepository: taskRepository,
	}
}

func (t *TaskServiceImpl) Create(request dto.CreateTaskRequest) (domain.TaskID, error) {
	task := domain.Task{
		ID:        domain.NewTaskId(),
		Name:      request.Name,
		ProjectID: request.ProjectID,
		Status:    domain.Todo,
	}
	err := t.taskRepository.Save(task)

	if err != nil {
		return domain.TaskID{}, err
	}
	return task.ID, nil
}

func (t *TaskServiceImpl) GetByID(id domain.TaskID) (dto.TaskDTO, error) {
	task, err := t.taskRepository.Get(id)
	return dto.ToTaskDTO(task), err
}

func (t *TaskServiceImpl) FindByProjectID(projectId domain.ProjectID) ([]dto.TaskDTO, error) {
	tasks := t.taskRepository.FindByProject(projectId)
	var dtos []dto.TaskDTO
	for _, task := range tasks {
		dtos = append(dtos, dto.ToTaskDTO(task))
	}
	return dtos, nil
}

func (t *TaskServiceImpl) DeleteByProjectID(projectId domain.ProjectID) error {
	tasks, err := t.FindByProjectID(projectId)
	if err != nil {
		return err
	}
	for _, task := range tasks {
		if task.ProjectID == projectId {
			if t.taskRepository.DeleteById(task.TaskID) != nil {
				return errors.New("failed to delete task with id: " + task.TaskID.String())
			}
		}
	}
	return nil
}

func (t *TaskServiceImpl) ChangeStatus(id domain.TaskID, status domain.TaskStatus) error {
	task, err := t.taskRepository.Get(id)
	if err != nil {
		return err
	}
	task.Status = status
	_, err = t.taskRepository.Update(task)
	return err
}
