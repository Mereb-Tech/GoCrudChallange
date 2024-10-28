package usecase

import (
	"fmt"
	Domain "mereb/Domain"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type PersonUsecase struct {
	personRepo      Domain.PersonRepository
	personValidator *validator.Validate
}

func NewPersonUsecase(personRepo Domain.PersonRepository) Domain.PersonUsecase {
	validate := validator.New()
	return &PersonUsecase{
		personRepo:      personRepo,
		personValidator: validate,
	}

}

func (usecase *PersonUsecase) GetAllPersons() ([]Domain.Person, error) {
	return usecase.personRepo.GetAllPersons()

}
func (usecase *PersonUsecase) GetPersonByID(id string) (Domain.Person, error) {
	return usecase.personRepo.GetPersonByID(id)
}
func (usecase *PersonUsecase) CreatePerson(person Domain.Person) (Domain.Person, error) {
	person.ID = uuid.New().String()
	fmt.Print(person)
	err := usecase.personValidator.Struct(person)
	if err != nil {
		return Domain.Person{}, err
	}
	_person, _err := usecase.personRepo.CreatePerson(person)
	return _person, _err
}
func (usecase *PersonUsecase) UpdatePerson(id string, person Domain.Person) (Domain.Person, error) {
	// check existance of person
	person, err := usecase.personRepo.GetPersonByID(id)
	if err != nil {
		return Domain.Person{}, err
	}
	person.ID = id
	err = usecase.personValidator.Struct(person)
	if err != nil {
		return Domain.Person{}, err
	}
	_person, _err := usecase.personRepo.UpdatePerson(id, person)
	return _person, _err
}
func (usecase *PersonUsecase) DeletePerson(id string) error {
	_, err := usecase.personRepo.GetPersonByID(id)
	if err != nil {
		return err
	}
	err = usecase.personRepo.DeletePerson(id)
	if err != nil {
		return err
	}
	return nil
}
