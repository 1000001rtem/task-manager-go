package dto

import "task-manager-go/internal/domain"

type CreateUserRequest struct {
	Login    string
	Email    string
	Password string
}

type UserDTO struct {
	ID    domain.UserID
	Login string
	Email string
}

func ToUserDTO(d domain.User) UserDTO {
	return UserDTO{
		ID:    d.ID,
		Login: d.Login,
		Email: d.Email,
	}
}
