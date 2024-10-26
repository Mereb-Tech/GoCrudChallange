package interfaces

import (
	dtos "GoCrudChallenge/Domain/DTOs"
	models "GoCrudChallenge/Domain/Models"
)



type PersonUseCaseInterface interface {
	CreatePerson(personRequest dtos.PersonRequestDTO) (*models.Person, error)
	GetPersonByID(id string) (*models.Person, error)
	UpdatePerson(id string, personRequest dtos.PersonRequestDTO) (*models.Person, error)
	DeletePerson(id string) error
	GetAllPersons() ([]*models.Person, error)
}