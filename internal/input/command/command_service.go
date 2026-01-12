package command

import (
	"errors"
	"fmt"
	"task-manager-go/internal/business/service/security"
	"task-manager-go/internal/domain"
)

type CommandService interface {
	Run(name string) error
}

type commandService struct {
	storage *Storage
}

func NewCommandService(storage *Storage) CommandService {
	return &commandService{
		storage: storage,
	}
}

func (service *commandService) Run(name string) error {
	command, ok := service.storage.Get(name)
	if !ok {
		return errors.New(fmt.Sprintf("command %s not found", name))
	}

	_, err := security.SecurityContext.GetCurrentUser()

	if err != nil && command.Secure == true {
		return domain.FormatError(domain.ErrAccessDenied, "command %s is secured and need authentication", name)
	}
	return command.Action()
}
