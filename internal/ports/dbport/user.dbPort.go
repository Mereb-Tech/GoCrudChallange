package dbport

import "github.com/Mahider-T/GoCrudChallange/internal/application/core/domain"

type UserDbPort interface {
	Create(user domain.User) (*domain.User, error)
	ReadOne(id string) (*domain.User, error)
	ReadAll() ([]*domain.User, error)
	Update(id string, user domain.User) (*domain.User, error)
	Delete(id string) error
}
