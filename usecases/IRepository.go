package usecases

import (
	"github.com/abe16s/GoCrudChallange/domain"
	"github.com/google/uuid"
)

type IRepo interface {
	Create(person domain.Person) error
	GetAllPersons() ([]*domain.Person, error)
	GetPersonById(id uuid.UUID) (*domain.Person, error)
	Update(person domain.Person) error
	Delete(id uuid.UUID) error
}