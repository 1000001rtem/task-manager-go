package command

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
	"task-manager-go/internal/business/dto"
	"task-manager-go/internal/business/service/project"
	"task-manager-go/internal/business/service/security"
	"task-manager-go/internal/business/service/task"
	"task-manager-go/internal/business/service/user"
)

type Storage struct {
	userService     user.UserService
	taskService     task.TaskService
	projectService  project.ProjectService
	securityService security.SecurityService
	reader          *bufio.Reader

	storage map[string]Command
}

func NewStorage(
	userService user.UserService,
	taskService task.TaskService,
	projectService project.ProjectService,
	securityService security.SecurityService,
	reader *bufio.Reader,
) *Storage {
	storage := &Storage{
		userService:     userService,
		taskService:     taskService,
		projectService:  projectService,
		securityService: securityService,
		reader:          reader,
	}
	storage.setUp()
	return storage
}

func (s *Storage) All() []Command {
	return slices.Collect(maps.Values(s.storage))
}

func (s *Storage) Get(name string) (Command, bool) {
	command, ok := s.storage[name]
	return command, ok
}

func (s *Storage) setUp() {
	s.storage = map[string]Command{
		"help": {
			"help",
			"Prints all commands",
			false,
			func() error {
				for _, command := range s.All() {
					fmt.Println(command.Name + " - " + command.Description)
				}
				return nil
			},
		},
		"login": {
			"login",
			"User authentication",
			false,
			func() error {
				fmt.Println("Please enter your login: ")
				login, loginErr := s.reader.ReadString('\n')
				fmt.Println("Please enter your password: ")
				password, passwordErr := s.reader.ReadString('\n')
				if loginErr != nil {
					return loginErr
				}
				if passwordErr != nil {
					return passwordErr
				}
				return s.securityService.Login(strings.TrimSpace(login), strings.TrimSpace(password))
			},
		},
		"project_create": {
			"project_create",
			"Create new project",
			true,
			func() error {
				fmt.Println("Please enter your project name: ")
				projectName, projectNameErr := s.reader.ReadString('\n')
				if projectNameErr != nil {
					return projectNameErr
				}
				currentUser, err := security.SecurityContext.GetCurrentUser()
				if err != nil {
					return err
				}
				_, err = s.projectService.Create(currentUser.ID, dto.CreateProjectRequest{Name: projectName})
				return err
			},
		},
		"project_all": {
			"project_all",
			"Returns all your projects",
			true,
			func() error {
				currentUser, err := security.SecurityContext.GetCurrentUser()
				if err != nil {
					return err
				}
				projects, err := s.projectService.FindByUserId(currentUser.ID)
				if err != nil {
					return err
				}
				for _, p := range projects {
					fmt.Println(p)
				}
				return nil
			},
		},
		"exit": {
			"exit",
			"Exit from application",
			false,
			func() error {
				fmt.Println("Bye Bye")
				os.Exit(0)
				return nil
			},
		},
	}
}
