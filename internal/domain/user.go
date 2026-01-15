package domain

import "github.com/google/uuid"

type UserID uuid.UUID

func (id UserID) String() string {
	return uuid.UUID(id).String()
}

func NewUserId() UserID {
	return UserID(uuid.New())
}

type User struct {
	ID       UserID
	Login    string
	Email    string
	Password string
}
