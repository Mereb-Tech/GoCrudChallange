package usecase

import (
	Domain "mereb/Domain"
	"github.com/go-playground/validator/v10"
)

type PersonUsecase struct {
	personRepo Domain.PersonRepository
	personValidator  *validator.Validate
}

func NewPersonUsecase(personRepo Domain.PersonRepository) Domain.PersonUsecase {
	validate := validator.New()
	return &PersonUsecase{
		personRepo: personRepo,
		personValidator: validate,
	}
}

func (usecase *PersonUsecase) GetAllPerson() ([]Domain.Person, error) {
	return []Domain.Person{}, nil
}
func (usecase *PersonUsecase) GetPersonByID(id string) (Domain.Person, error) {
	return Domain.Person{}, nil
}
func (usecase *PersonUsecase) CreatePerson(person Domain.Person) (Domain.Person, error) {
	return Domain.Person{}, nil
}
func (usecase *PersonUsecase) UpdatePerson(id string, person Domain.Person) (Domain.Person, error) {
	return Domain.Person{}, nil
}
func (usecase *PersonUsecase) DeletePerson(id string) (Domain.Person, error) {
	return Domain.Person{}, nil
}

