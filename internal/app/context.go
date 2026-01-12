package app

import (
	"bufio"
	"os"
	"task-manager-go/internal/business/service/project"
	"task-manager-go/internal/business/service/security"
	"task-manager-go/internal/business/service/task"
	"task-manager-go/internal/business/service/user"
	"task-manager-go/internal/input/command"
	"task-manager-go/internal/repository/memory"
)

type appContext struct {
	UserService    user.UserService
	TaskService    task.TaskService
	ProjectService project.ProjectService
	AuthService    security.SecurityService
	CommandService command.CommandService

	Reader *bufio.Reader
}

var Context = createContext()

func createContext() *appContext {
	store := memory.NewStore()

	userService := user.NewUserService(store.UserRepository)
	taskService := task.NewTaskService(store.TaskRepository)
	projectService := project.NewProjectService(store.ProjectRepository, taskService)
	securityService := security.NewSecurityService(store.UserRepository)

	reader := bufio.NewReader(os.Stdin)
	commandStorage := command.NewStorage(
		userService,
		taskService,
		projectService,
		securityService,
		reader,
	)
	commandService := command.NewCommandService(commandStorage)

	context := &appContext{
		UserService:    userService,
		TaskService:    taskService,
		ProjectService: projectService,
		AuthService:    securityService,
		CommandService: commandService,
		Reader:         reader,
	}
	context.CommandService = commandService

	return context
}
