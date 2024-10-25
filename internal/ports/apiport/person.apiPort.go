package apiport

import "github.com/Mahider-T/GoCrudChallange/internal/application/core/domain"

type PersonApiPort interface {
	CreatePerson(person domain.CreatePersonDTO) (*domain.Person, error)
	GetPersonByID(id string) (*domain.Person, error)
	GetAllPersons() ([]*domain.Person, error)
	UpdatePerson(id string, person domain.UpdatePersonDTO) (*domain.Person, error)
	DeletePerson(id string) error
}
