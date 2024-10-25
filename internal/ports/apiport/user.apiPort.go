package apiport

import "github.com/Mahider-T/GoCrudChallange/internal/application/core/domain"

type UserApiPort interface {
	CreateUser(user domain.User) (*domain.User, error)
	GetUserByID(id string) (*domain.User, error)
	GetAllUsers() ([]*domain.User, error)
	UpdateUser(id string, user domain.UpdateUserDTO) (*domain.User, error)
	DeleteUser(id string) error
}
