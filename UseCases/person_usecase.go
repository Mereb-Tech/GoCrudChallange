package usecases

import (
	dtos "GoCrudChallenge/Domain/DTOs"
	interfaces "GoCrudChallenge/Domain/Interfaces"
	models "GoCrudChallenge/Domain/Models"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type PersonUseCase struct {
	Repo      interfaces.PersonRepository
	Validator *validator.Validate
}

func NewPersonUseCase(Repo interfaces.PersonRepository) interfaces.PersonUseCaseInterface {
	return &PersonUseCase{
		Repo:      Repo,
		Validator: validator.New(),
	}
}

func (uc *PersonUseCase) CreatePerson(personRequest dtos.PersonRequestDTO) (*models.Person, error) {
	person := &models.Person{
		ID:      uuid.NewString(),
		Name:    personRequest.Name,
		Age:     personRequest.Age,
		Hobbies: personRequest.Hobbies,
	}

	if err := uc.Validator.Struct(person); err != nil {
		return nil, fmt.Errorf("Invalid person data")
	}

	if err := uc.Repo.CreatePerson(person); err != nil {
		return nil, err
	}

	return person, nil
}

func (uc *PersonUseCase) GetPersonByID(id string) (*models.Person, error) {
	person, err := uc.Repo.GetPersonByID(id)
	if err != nil {
		return nil, err
	}
	return person, nil
}

func (uc *PersonUseCase) UpdatePerson(id string, personRequest dtos.PersonRequestDTO) (*models.Person, error) {
	person := &models.Person{
		ID:      id,
		Name:    personRequest.Name,
		Age:     personRequest.Age,
		Hobbies: personRequest.Hobbies,
	}

	if err := uc.Repo.UpdatePerson(person); err != nil {
		return nil, err
	}

	updatedPerson, err := uc.Repo.GetPersonByID(id)
	if err != nil {
		return nil, err
	}

	return updatedPerson, nil
}

func (uc *PersonUseCase) DeletePerson(id string) error {
	err := uc.Repo.DeletePerson(id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *PersonUseCase) GetAllPersons() ([]*models.Person, error) {
	persons := uc.Repo.GetAllPersons()
	if persons == nil {
		return nil, fmt.Errorf("No persons found")
	}

	return persons, nil
}
