package api

import (
	"github.com/Mahider-T/GoCrudChallange/internal/application/core/domain"
	"github.com/Mahider-T/GoCrudChallange/internal/ports/dbport"
)

type PersonApi struct {
	PersonDbPort dbport.PersonDbPort
}

func NewPersonApi(personDbPort dbport.PersonDbPort) *PersonApi {
	return &PersonApi{
		PersonDbPort: personDbPort,
	}
}

func (ua *PersonApi) CreatePerson(person domain.CreatePersonDTO) (*domain.Person, error) {
	if err := domain.ValidateCreatePersonDTO(person); err != nil {
		return nil, domain.ErrBadRequest
	}
	return ua.PersonDbPort.Create(person)

}

func (ua *PersonApi) GetPersonByID(id string) (*domain.Person, error) {
	return ua.PersonDbPort.ReadOne(id)
}

func (ua *PersonApi) GetAllPersons() ([]*domain.Person, error) {
	return ua.PersonDbPort.ReadAll()
}

func (ua *PersonApi) UpdatePerson(id string, person domain.UpdatePersonDTO) (*domain.Person, error) {

	if err := domain.ValidateUpdatePersonDTO(person); err != nil {
		return nil, domain.ErrBadRequest
	}

	existingPerson, err := ua.GetPersonByID(id)
	if err != nil {
		return nil, err
	}
	if person.Age != nil {
		existingPerson.Age = *person.Age
	}
	if person.Name != nil {
		existingPerson.Name = *person.Name
	}
	if person.Hobbies != nil {
		existingPerson.Hobbies = *person.Hobbies
	}

	return ua.PersonDbPort.Update(id, *existingPerson)
}

func (ua *PersonApi) DeletePerson(id string) error {
	return ua.PersonDbPort.Delete(id)
}
